package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVendorResponse Response Object
type UpdateVendorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVendorResponse struct{}"
	}

	return strings.Join([]string{"UpdateVendorResponse", string(data)}, " ")
}
