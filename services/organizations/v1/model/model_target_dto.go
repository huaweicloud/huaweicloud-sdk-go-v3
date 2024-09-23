package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetDto 要邀请加入组织的账号的标识符（ID）。
type TargetDto struct {

	// 目标类型，account：账户id，name：账户名称。
	Type string `json:"type"`

	// 如果指定 'type:account'，则必须提供账号ID作为实体。如果指定 'type:name'，则必须指定账号名称作为实体。
	Entity string `json:"entity"`
}

func (o TargetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDto struct{}"
	}

	return strings.Join([]string{"TargetDto", string(data)}, " ")
}
