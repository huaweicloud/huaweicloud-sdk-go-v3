package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttackTotalResponse Response Object
type ShowAttackTotalResponse struct {
	Data           *ShowAttackTotalRespData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAttackTotalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTotalResponse struct{}"
	}

	return strings.Join([]string{"ShowAttackTotalResponse", string(data)}, " ")
}
