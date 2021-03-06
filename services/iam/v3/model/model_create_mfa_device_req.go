package model

import (
	"encoding/json"

	"strings"
)

//
type CreateMfaDeviceReq struct {
	VirtualMfaDevice *CreateMfaDevice `json:"virtual_mfa_device"`
}

func (o CreateMfaDeviceReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMfaDeviceReq struct{}"
	}

	return strings.Join([]string{"CreateMfaDeviceReq", string(data)}, " ")
}
