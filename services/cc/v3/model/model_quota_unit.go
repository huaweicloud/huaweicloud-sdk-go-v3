package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaUnit struct {

	// 配额值的单位。
	Unit string `json:"unit"`
}

func (o QuotaUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaUnit struct{}"
	}

	return strings.Join([]string{"QuotaUnit", string(data)}, " ")
}
