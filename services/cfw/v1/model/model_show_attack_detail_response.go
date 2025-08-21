package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttackDetailResponse Response Object
type ShowAttackDetailResponse struct {
	Data           *AttackDetailVo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAttackDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAttackDetailResponse", string(data)}, " ")
}
