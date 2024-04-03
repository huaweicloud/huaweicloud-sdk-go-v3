package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaList struct {
	Resources *[]QuotaInfo `json:"resources,omitempty"`
}

func (o QuotaList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaList struct{}"
	}

	return strings.Join([]string{"QuotaList", string(data)}, " ")
}
