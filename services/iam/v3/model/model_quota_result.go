package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type QuotaResult struct {
	// 资源信息

	Resources *[]Resources `json:"resources,omitempty"`
}

func (o QuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResult struct{}"
	}

	return strings.Join([]string{"QuotaResult", string(data)}, " ")
}
