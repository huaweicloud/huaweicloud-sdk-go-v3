package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateDnatOptionBody 创建DNAT规则的请求体。
type CreatePrivateDnatOptionBody struct {
	DnatRule *CreatePrivateDnatOption `json:"dnat_rule"`
}

func (o CreatePrivateDnatOptionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateDnatOptionBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateDnatOptionBody", string(data)}, " ")
}
