package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelVendorResponse Response Object
type DeleteModelVendorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteModelVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelVendorResponse struct{}"
	}

	return strings.Join([]string{"DeleteModelVendorResponse", string(data)}, " ")
}
