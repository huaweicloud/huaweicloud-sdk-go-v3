package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityPolicyControlResourceItemVo 安全策略管控资源列表项响应。
type SecurityPolicyControlResourceItemVo struct {

	// 资源记录ID。
	Id *string `json:"id,omitempty"`

	// 资源类型，DESKTOP或DESKTOP_TAG。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源ID（DESKTOP时为instance_id，DESKTOP_TAG时为key:value）。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称（DESKTOP时为实例名称，DESKTOP_TAG时为标签）。
	ResourceName *string `json:"resource_name,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o SecurityPolicyControlResourceItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityPolicyControlResourceItemVo struct{}"
	}

	return strings.Join([]string{"SecurityPolicyControlResourceItemVo", string(data)}, " ")
}
