package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModuleShadowResponse Response Object
type ShowModuleShadowResponse struct {

	// 应用配置内容
	Properties *interface{} `json:"properties,omitempty"`

	// 应用配置更新时间
	PropertiesUpdateTime *interface{} `json:"properties_update_time,omitempty"`
	HttpStatusCode       int          `json:"-"`
}

func (o ShowModuleShadowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModuleShadowResponse struct{}"
	}

	return strings.Join([]string{"ShowModuleShadowResponse", string(data)}, " ")
}
