package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {
	Instance *InstancesVo `json:"instance,omitempty" xml:"instance"`

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 静态资源链接
	BundleUrl      *string `json:"bundle_url,omitempty" xml:"bundle_url"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
