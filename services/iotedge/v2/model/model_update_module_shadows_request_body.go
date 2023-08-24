package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateModuleShadowsRequestBody struct {

	// 应用配置内容
	Properties *interface{} `json:"properties,omitempty"`
}

func (o UpdateModuleShadowsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModuleShadowsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateModuleShadowsRequestBody", string(data)}, " ")
}
