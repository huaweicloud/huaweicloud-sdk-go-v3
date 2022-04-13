package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type QuotaResources struct {
	// 配额信息列表。

	Resources *[]QuotaResource `json:"resources,omitempty"`
}

func (o QuotaResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResources struct{}"
	}

	return strings.Join([]string{"QuotaResources", string(data)}, " ")
}
