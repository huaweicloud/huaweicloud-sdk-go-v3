package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchListResourceResponseData struct {

	// CMDB分配的资源ID
	Id *string `json:"id,omitempty"`

	// 云服务分配的资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// Openstack中的项目I
	ProjectId *string `json:"project_id,omitempty"`

	// region ID
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 资源标签
	Tags *[]Tag `json:"tags,omitempty"`

	// uniAgent唯一id
	AgentId *string `json:"agent_id,omitempty"`

	// uniAgent状态
	AgentState *string `json:"agent_state,omitempty"`

	// 资源详细属性
	Properties map[string]interface{} `json:"properties,omitempty"`

	// 采集属性
	IngestProperties map[string]string `json:"ingest_properties,omitempty"`

	// 是否已托管
	IsDelegated *bool `json:"is_delegated,omitempty"`
}

func (o BatchListResourceResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListResourceResponseData struct{}"
	}

	return strings.Join([]string{"BatchListResourceResponseData", string(data)}, " ")
}
