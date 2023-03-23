package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新SNAT规则的请求体。
type UpdatePrivateSnatOptionBody struct {
	SnatRule *UpdatePrivateSnatOption `json:"snat_rule"`
}

func (o UpdatePrivateSnatOptionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatOptionBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatOptionBody", string(data)}, " ")
}
