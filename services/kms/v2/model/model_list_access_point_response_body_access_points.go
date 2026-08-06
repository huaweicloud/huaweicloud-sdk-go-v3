package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAccessPointResponseBodyAccessPoints struct {

	// **参数解释：** 接入点ID **取值范围：** 不涉及
	AccessPointId string `json:"access_point_id"`

	// **参数解释：** 密钥空间ID **取值范围：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 接入点名称 **取值范围：** 不涉及
	AccessPointName string `json:"access_point_name"`

	// **参数解释：** 接入点状态 **取值范围：** 0:禁用，1：启用
	State int32 `json:"state"`

	// **参数解释：** 接入点类型 **取值范围：** 1:ECS，2：CCE，3：Custom
	Type int32 `json:"type"`

	// **参数解释：** 接入点创建人 **取值范围：** 不涉及
	CreatedBy string `json:"created_by"`

	// **参数解释：** 接入点创建时间 **取值范围：** 不涉及
	CreateTime string `json:"create_time"`

	// **参数解释：** 接入点最近更新时间 **取值范围：** 不涉及
	LsatModifyTime string `json:"lsat_modify_time"`
}

func (o ListAccessPointResponseBodyAccessPoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPointResponseBodyAccessPoints struct{}"
	}

	return strings.Join([]string{"ListAccessPointResponseBodyAccessPoints", string(data)}, " ")
}
