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

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Type *ConnectionPointTypeEnum `json:"type"`

	// 实例ID。
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`
}

func (o ConnectionPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPoint struct{}"
	}

	return strings.Join([]string{"ConnectionPoint", string(data)}, " ")
}
