package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/bensallen/rbd/models"
	"github.com/bensallen/rbd/pkg/mount"
)

type MountsOverlayType struct {
	last  int64
	mnts  map[int64]*models.MountOverlay
	mutex *sync.Mutex
}

func (m *MountsOverlayType) Init() {
	m.last = 0
	m.mnts = make(map[int64]*models.MountOverlay)
	m.mutex = &sync.Mutex{}
}

func (m *MountsOverlayType) List() (r []*models.MountOverlay) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, mnt := range m.mnts {
		r = append(r, mnt)
	}
	return
}

func (m *MountsOverlayType) Get(id int64) (*models.MountOverlay, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if r, ok := m.mnts[id]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("no overlay mount by id %d", id)
}

func (m *MountsOverlayType) Mount(mnt *models.MountOverlay) (r *models.MountOverlay, err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// note: we don't check for existence, because we will happily make more than one overlay of the same thing

	// there most be at least one lower
	if len(mnt.Lower) == 0 {
		return nil, fmt.Errorf("at least one lower mount must be specified")
	}

	// make sure the lower mounts exists/are ours + assemble list of lower mountpoints
	// warning: there's a possible race here if someone removed these mounts while we're assembling
	//			we might need an extneral interface to lock them.
	lmnts := []string{}
	for _, i := range mnt.Lower {
		var rmnt *models.MountRbd
		if rmnt, err = MountsRbd.Get(i); err != nil {
			return nil, fmt.Errorf("mount failure: %v", err)
		}
		lmnts = append(lmnts, rmnt.Mountpoint)
	}

	// ok, we're good to attempt the mount
	// make a mountpoint/upperdir/workdir
	if err = os.MkdirAll(mountDir, 0700); err != nil {
		return nil, fmt.Errorf("could not create base mount directory: %v", err)
	}
	if mnt.Mountpoint, err = ioutil.TempDir(mountDir, "mount_"); err != nil {
		return nil, fmt.Errorf("could not create mountpoint: %v", err)
	}
	if mnt.Upperdir, err = ioutil.TempDir(mountDir, "upper_"); err != nil {
		return nil, fmt.Errorf("could not create upperdir: %v", err)
	}
	if mnt.Workdir, err = ioutil.TempDir(mountDir, "work_"); err != nil {
		return nil, fmt.Errorf("could not create workdir: %v", err)
	}

	// try the mounmt
	opts := []string{
		"lowerdir=" + strings.Join(lmnts, ":"),
		"upperdir=" + mnt.Upperdir,
		"workdir=" + mnt.Workdir,
	}
	if err = mount.Mount("overlay", mnt.Mountpoint, "overlay", opts); err != nil {
		return nil, fmt.Errorf("mount failure: %v", err)
	}

	// store
	m.last++
	mnt.ID = m.last
	m.mnts[mnt.ID] = mnt

	// add refs
	for _, i := range mnt.Lower {
		MountsRbd.RefAdd(i, 1)
	}
	return mnt, nil
}

func (m *MountsOverlayType) Unmount(id int64) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var mnt *models.MountOverlay
	var ok bool

	if mnt, ok = m.mnts[id]; !ok {
		return fmt.Errorf("unmount failure: no such device %d", id)
	}

	if mnt.Ref > 0 {
		return fmt.Errorf("unmount failure: mount is in use")
	}

	// always lazy unmount.  Good idea?
	if err = mount.Unmount(mnt.Mountpoint, false, true); err != nil {
		return fmt.Errorf("unmount failure: %v", err)
	}

	os.Remove(mnt.Mountpoint)  // we shouldn't fail on this. Should we report it anyway?
	os.RemoveAll(mnt.Workdir)  // option to leave behind?
	os.RemoveAll(mnt.Upperdir) // option to leave behind? Or store on RBD?
	delete(m.mnts, id)
	for _, i := range mnt.Lower {
		MountsRbd.RefAdd(i, -1)
	}
	return
}

func (m *MountsOverlayType) RefAdd(id, n int64) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if r, ok := m.mnts[id]; ok {
		r.Ref += n
	}
}

func (*MountsOverlayType) refAddList(mnts []*models.MountRbd, n int64) {
	for _, mnt := range mnts {
		MountsRbd.RefAdd(*mnt.ID, n)
	}
}
