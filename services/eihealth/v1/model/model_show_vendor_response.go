package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVendorResponse struct {

	// 商标Base64
	Logo *string `json:"logo,omitempty"`

	// 名称
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVendorResponse struct{}"
	}

	return strings.Join([]string{"ShowVendorResponse", string(data)}, " ")
}
