package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateSnatOptionBody 创建SNAT规则的请求体。
type CreatePrivateSnatOptionBody struct {
	SnatRule *CreatePrivateSnatOption `json:"snat_rule"`
}

func (o CreatePrivateSnatOptionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatOptionBody struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatOptionBody", string(data)}, " ")
}
