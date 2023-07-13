package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppInstanceRequestDto struct {

	// 应用实例ID
	Id string `json:"id"`

	// 边缘集群命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	AppVersion string `json:"app_version"`

	// 应用实例chart配置
	Values *interface{} `json:"values,omitempty"`
}

func (o CreateAppInstanceRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppInstanceRequestDto struct{}"
	}

	return strings.Join([]string{"CreateAppInstanceRequestDto", string(data)}, " ")
}
