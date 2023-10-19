package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionPoint 中心网络连接点定义。
type ConnectionPoint struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 站点编码定义
	SiteCode string `json:"site_code"`

	// 资源ID标识符。
	InstanceId string `json:"instance_id"`

	Type *ConnectionPointTypeEnum `json:"type"`
}

func (o ConnectionPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPoint struct{}"
	}

	return strings.Join([]string{"ConnectionPoint", string(data)}, " ")
}
