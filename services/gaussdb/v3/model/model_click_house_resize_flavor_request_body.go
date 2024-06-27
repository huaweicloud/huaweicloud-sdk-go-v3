package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseResizeFlavorRequestBody 规格变更请求体。
type ClickHouseResizeFlavorRequestBody struct {

	// 规格ID。可通过“查询规格信息”接口获取。  仅允许使用对应数据库，和对应实例类型的规格ID。
	FlavorRef string `json:"flavor_ref"`

	// 是否延迟变更。默认false。
	Delay bool `json:"delay"`

	// 实例ID，严格匹配UUID规则。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ClickHouseResizeFlavorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseResizeFlavorRequestBody struct{}"
	}

	return strings.Join([]string{"ClickHouseResizeFlavorRequestBody", string(data)}, " ")
}
