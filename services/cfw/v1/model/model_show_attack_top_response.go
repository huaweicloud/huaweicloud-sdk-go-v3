package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttackTopResponse Response Object
type ShowAttackTopResponse struct {
	Data           *AttackTopRespBody `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowAttackTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTopResponse struct{}"
	}

	return strings.Join([]string{"ShowAttackTopResponse", string(data)}, " ")
}
