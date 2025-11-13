package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelVendorResponse Response Object
type UpdateModelVendorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateModelVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelVendorResponse struct{}"
	}

	return strings.Join([]string{"UpdateModelVendorResponse", string(data)}, " ")
}
