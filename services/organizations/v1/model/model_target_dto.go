package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 要邀请加入组织的帐号的标识符（ID）。
type TargetDto struct {

	// 目标类型，account：账户，email：邮箱。
	Type string `json:"type"`

	// 如果指定 'type:account'，则必须提供帐号ID作为实体。如果指定'type:email'，则必须指定与帐号关联的电子邮件地址。
	Entity string `json:"entity"`
}

func (o TargetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDto struct{}"
	}

	return strings.Join([]string{"TargetDto", string(data)}, " ")
}
