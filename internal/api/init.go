package api

var Rbds RbdsType
var MountsRbd MountsRBDType

func init() {
	Rbds = RbdsType{}
	Rbds.Init()
	MountsRbd = MountsRBDType{}
	MountsRbd.Init()
}
