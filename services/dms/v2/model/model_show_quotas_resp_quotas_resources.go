package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowQuotasRespQuotasResources struct {

	// 配额名称。
	Type *string `json:"type,omitempty" xml:"type"`

	// 配额数量。
	Quota *int32 `json:"quota,omitempty" xml:"quota"`

	// 已使用的数量。
	Used *int32 `json:"used,omitempty" xml:"used"`

	// 配额调整的最小值。
	Min *int32 `json:"min,omitempty" xml:"min"`

	// 配额调整的最大值。
	Max *int32 `json:"max,omitempty" xml:"max"`
}

func (o ShowQuotasRespQuotasResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRespQuotasResources struct{}"
	}

	return strings.Join([]string{"ShowQuotasRespQuotasResources", string(data)}, " ")
}
