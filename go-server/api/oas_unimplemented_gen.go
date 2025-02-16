// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DeleteV1DiningUser implements deleteV1DiningUser operation.
//
// Deletes a user given a session.
//
// DELETE /v1/dining/user
func (UnimplementedHandler) DeleteV1DiningUser(ctx context.Context, req OptDiningUserSession) (r DeleteV1DiningUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1Alerts implements getV1Alerts operation.
//
// Returns all of BearTrak's active alerts.
//
// GET /v1/alerts
func (UnimplementedHandler) GetV1Alerts(ctx context.Context) (r []Alert, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1Dining implements getV1Dining operation.
//
// Returns all necessary data for BearTrak's dining section.
//
// GET /v1/dining
func (UnimplementedHandler) GetV1Dining(ctx context.Context) (r []Eatery, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1DiningUserAccounts implements getV1DiningUserAccounts operation.
//
// Returns a dining user's transaction accounts given a session.
//
// GET /v1/dining/user/accounts
func (UnimplementedHandler) GetV1DiningUserAccounts(ctx context.Context, req OptDiningUserSession) (r GetV1DiningUserAccountsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1DiningUserBarcode implements getV1DiningUserBarcode operation.
//
// Returns a user's dining hall barcode given a session.
//
// GET /v1/dining/user/barcode
func (UnimplementedHandler) GetV1DiningUserBarcode(ctx context.Context, req OptDiningUserSession) (r GetV1DiningUserBarcodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1DiningUserSession implements getV1DiningUserSession operation.
//
// Refreshes a session given a user device.
//
// GET /v1/dining/user/session
func (UnimplementedHandler) GetV1DiningUserSession(ctx context.Context, req OptDiningUserDevice) (r GetV1DiningUserSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1Gyms implements getV1Gyms operation.
//
// Returns all necessary data for BearTrak's gym section.
//
// GET /v1/gyms
func (UnimplementedHandler) GetV1Gyms(ctx context.Context) (r []Gym, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1TransitRoutes implements getV1TransitRoutes operation.
//
// Returns non time-sensitive, route-related data for BearTrak's transit section.
//
// GET /v1/transit/routes
func (UnimplementedHandler) GetV1TransitRoutes(ctx context.Context) (r []BusRoute, _ error) {
	return r, ht.ErrNotImplemented
}

// GetV1TransitVehicles implements getV1TransitVehicles operation.
//
// Returns time-sensitive, vehicle-related data for BearTrak's transit section.
//
// GET /v1/transit/vehicles
func (UnimplementedHandler) GetV1TransitVehicles(ctx context.Context) (r []Vehicle, _ error) {
	return r, ht.ErrNotImplemented
}

// PostV1DiningUser implements postV1DiningUser operation.
//
// Registers a new user given a device and session.
//
// POST /v1/dining/user
func (UnimplementedHandler) PostV1DiningUser(ctx context.Context, req OptPostV1DiningUserReq) (r PostV1DiningUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
