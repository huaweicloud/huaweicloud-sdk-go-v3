package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库实例所在Region，项目等信息，当数据库实例类型为ecs（华为云ECS自建数据库），cloud（华为云数据库）时为必填项。
type CloudBaseInfo struct {

	// 区域ID，当数据库实例类型为ecs（华为云ECS自建数据库），cloud（华为云数据库）时为必填项。获取方法请参见地区和终端节点。 注意：当该Region下存在子项目时，Region ID为区域项目ID与子项目ID，由“_”下划线拼接，例如：cn-north-4_abc。
	Region string `json:"region"`

	// 租户在某一Region下的Project ID。 获取方法请参见获取项目ID。
	ProjectId string `json:"project_id"`

	// 数据库所在可用分区（AZ）名称。
	AzCode *string `json:"az_code,omitempty"`
}

func (o CloudBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudBaseInfo struct{}"
	}

	return strings.Join([]string{"CloudBaseInfo", string(data)}, " ")
}
