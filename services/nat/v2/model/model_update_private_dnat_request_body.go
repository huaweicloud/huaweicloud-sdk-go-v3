package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateDnatRequestBody 更新DNAT规则的请求体。
type UpdatePrivateDnatRequestBody struct {
	DnatRule *UpdatePrivateDnatOption `json:"dnat_rule,omitempty"`
}

func (o UpdatePrivateDnatRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateDnatRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivateDnatRequestBody", string(data)}, " ")
}
