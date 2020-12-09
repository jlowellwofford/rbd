// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	internal "github.com/bensallen/rbd/internal/api"
	"github.com/bensallen/rbd/models"
	"github.com/bensallen/rbd/restapi/operations"
	"github.com/bensallen/rbd/restapi/operations/mounts"
	"github.com/bensallen/rbd/restapi/operations/rbds"
)

//go:generate swagger generate server --target ../../rbd --name Rbd --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.RbdAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RbdAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.RbdsListRbdsHandler = rbds.ListRbdsHandlerFunc(func(params rbds.ListRbdsParams) middleware.Responder {
		return rbds.NewListRbdsOK().WithPayload(internal.Rbds.List())
	})

	api.RbdsMapRbdHandler = rbds.MapRbdHandlerFunc(func(params rbds.MapRbdParams) middleware.Responder {
		var err error
		var r *models.Rbd
		if r, err = internal.Rbds.Map(params.Rbd); err != nil {
			return rbds.NewMapRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return rbds.NewMapRbdCreated().WithPayload(r)
	})

	api.RbdsGetRbdHandler = rbds.GetRbdHandlerFunc(func(params rbds.GetRbdParams) middleware.Responder {
		var err error
		var r *models.Rbd
		if r, err = internal.Rbds.Get(params.ID); err != nil {
			return rbds.NewGetRbdDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return rbds.NewGetRbdOK().WithPayload(r)
	})

	api.RbdsUnmapRbdHandler = rbds.UnmapRbdHandlerFunc(func(params rbds.UnmapRbdParams) middleware.Responder {
		if err := internal.Rbds.Unmap(params.ID); err != nil {
			return rbds.NewUnmapRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return rbds.NewUnmapRbdNoContent()
	})

	// MountsRbd

	api.MountsListMountsRbdHandler = mounts.ListMountsRbdHandlerFunc(func(params mounts.ListMountsRbdParams) middleware.Responder {
		return mounts.NewListMountsRbdOK().WithPayload(internal.MountsRbd.List())
	})

	api.MountsMountRbdHandler = mounts.MountRbdHandlerFunc(func(params mounts.MountRbdParams) middleware.Responder {
		var err error
		var r *models.MountRbd
		if r, err = internal.MountsRbd.Mount(params.Mount); err != nil {
			return mounts.NewMountRbdDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return mounts.NewMountRbdCreated().WithPayload(r)
	})

	api.MountsGetMountRbdHandler = mounts.GetMountRbdHandlerFunc(func(params mounts.GetMountRbdParams) middleware.Responder {
		var err error
		var r *models.MountRbd
		if r, err = internal.MountsRbd.Get(params.ID); err != nil {
			return mounts.NewGetMountRbdDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewGetMountRbdOK().WithPayload(r)
	})

	api.MountsUnmountRbdHandler = mounts.UnmountRbdHandlerFunc(func(params mounts.UnmountRbdParams) middleware.Responder {
		if err := internal.MountsRbd.Unmount(params.ID); err != nil {
			return mounts.NewUnmountRbdDefault(500).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewUnmountRbdNoContent()
	})

	// MountsOverlay

	api.MountsListMountsOverlayHandler = mounts.ListMountsOverlayHandlerFunc(func(params mounts.ListMountsOverlayParams) middleware.Responder {
		return mounts.NewListMountsOverlayOK().WithPayload(internal.MountsOverlay.List())
	})

	api.MountsMountOverlayHandler = mounts.MountOverlayHandlerFunc(func(params mounts.MountOverlayParams) middleware.Responder {
		var err error
		var r *models.MountOverlay
		if r, err = internal.MountsOverlay.Mount(params.Mount); err != nil {
			return mounts.NewMountOverlayDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return mounts.NewMountOverlayCreated().WithPayload(r)
	})

	api.MountsGetMountOverlayHandler = mounts.GetMountOverlayHandlerFunc(func(params mounts.GetMountOverlayParams) middleware.Responder {
		var err error
		var r *models.MountOverlay
		if r, err = internal.MountsOverlay.Get(params.ID); err != nil {
			return mounts.NewGetMountOverlayDefault(404).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewGetMountOverlayOK().WithPayload(r)
	})

	api.MountsUnmountOverlayHandler = mounts.UnmountOverlayHandlerFunc(func(params mounts.UnmountOverlayParams) middleware.Responder {
		if err := internal.MountsOverlay.Unmount(params.ID); err != nil {
			return mounts.NewUnmountOverlayDefault(500).WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return mounts.NewUnmountOverlayNoContent()
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
