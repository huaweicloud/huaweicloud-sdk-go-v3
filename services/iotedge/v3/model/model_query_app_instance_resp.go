package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryAppInstanceResp struct {

	// 应用实例ID
	Id *string `json:"id,omitempty"`

	// 边缘集群命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 应用实例版本
	Version *string `json:"version,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 应用实例状态
	Status *string `json:"status,omitempty"`

	// 状态描述
	StatusDescription *string `json:"status_description,omitempty"`

	// 应用实例chart配置
	Values *interface{} `json:"values,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryAppInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAppInstanceResp struct{}"
	}

	return strings.Join([]string{"QueryAppInstanceResp", string(data)}, " ")
}
