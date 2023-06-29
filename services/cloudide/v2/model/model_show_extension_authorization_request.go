package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionAuthorizationRequest Request Object
type ShowExtensionAuthorizationRequest struct {

	// 插件版本
	ExtensionVersion string `json:"extension_version"`

	// 插件标识(发布者.插件名)
	Identifier string `json:"identifier"`

	// CodeArtsIDEOnline实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowExtensionAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"ShowExtensionAuthorizationRequest", string(data)}, " ")
}
