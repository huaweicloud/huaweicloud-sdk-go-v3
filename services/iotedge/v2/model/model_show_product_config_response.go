package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProductConfigResponse struct {

	// 服务id
	ServiceId *string `json:"service_id,omitempty" xml:"service_id"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 属性
	Properties     *[]interface{} `json:"properties,omitempty" xml:"properties"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowProductConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowProductConfigResponse", string(data)}, " ")
}
