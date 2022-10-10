package msrandom

import (
	"github.com/com-gft-tsbo-source/go-common/ms-framework/microservice"
)

// ###########################################################################
// ###########################################################################
// MsRandom Response - Device
// ###########################################################################
// ###########################################################################

// RandomResponse ...
type RandomResponse struct {
	microservice.Response
	Value int `json:"value"`
}

// ###########################################################################

// InitRandomResponse Constructor of a response of ms-random
func InitRandomResponse(r *RandomResponse, code int, status string, ms *MsRandom) {
	microservice.InitResponseFromMicroService(&r.Response, ms, code, status)
	r.Value = 0
}

// NewRandomResponse ...
func NewRandomResponse(code int, status string, ms *MsRandom) *RandomResponse {
	var r RandomResponse
	InitRandomResponse(&r, code, status, ms)
	return &r
}
