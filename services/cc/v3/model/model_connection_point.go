package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionPoint 中心网络连接点定义。
type ConnectionPoint struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 站点编码定义
	SiteCode string `json:"site_code"`

	// 连接点的实例ID。
	InstanceId string `json:"instance_id"`

	// 连接点的实例的父资源ID。
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`

	Type *ConnectionPointTypeEnum `json:"type"`
}

func (o ConnectionPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPoint struct{}"
	}

	return strings.Join([]string{"ConnectionPoint", string(data)}, " ")
}
