package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserQuotaInfo struct {

	// 子账户（用户）ID。
	UserId string `json:"user_id"`

	Quota *ListUserQuotaDetail `json:"quota"`

	Usage *ListUserQuotaDetail `json:"usage,omitempty"`
}

func (o UserQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserQuotaInfo struct{}"
	}

	return strings.Join([]string{"UserQuotaInfo", string(data)}, " ")
}
