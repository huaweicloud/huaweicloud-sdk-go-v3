package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaRespQuotasResources struct {

	// **参数解释**： 类型 **取值范围**： - Charts：配额类型为模板
	Type *string `json:"type,omitempty"`

	// 配额
	Quota *int32 `json:"quota,omitempty"`

	// 已使用量
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaRespQuotasResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRespQuotasResources struct{}"
	}

	return strings.Join([]string{"QuotaRespQuotasResources", string(data)}, " ")
}
