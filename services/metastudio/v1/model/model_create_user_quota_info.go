package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUserQuotaInfo struct {

	// 子账户（用户）ID。
	UserId string `json:"user_id"`

	Quota *UserQuotaDetail `json:"quota"`
}

func (o CreateUserQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserQuotaInfo struct{}"
	}

	return strings.Join([]string{"CreateUserQuotaInfo", string(data)}, " ")
}
