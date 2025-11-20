package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateUserQuotaInfo struct {
	Quota *UserQuotaDetail `json:"quota"`
}

func (o UpdateUserQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserQuotaInfo struct{}"
	}

	return strings.Join([]string{"UpdateUserQuotaInfo", string(data)}, " ")
}
