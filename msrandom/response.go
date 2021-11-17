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
func InitRandomResponse(r *RandomResponse, status string, ms *MsRandom) {
	microservice.InitResponseFromMicroService(&r.Response, ms, status)
	r.Value = 0
}

// NewRandomResponse ...
func NewRandomResponse(status string, ms *MsRandom) *RandomResponse {
	var r RandomResponse
	InitRandomResponse(&r, status, ms)
	return &r
}
