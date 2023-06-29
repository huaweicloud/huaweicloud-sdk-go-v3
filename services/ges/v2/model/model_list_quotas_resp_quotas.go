package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotasRespQuotas resource类型列表，请求失败时该字段为空。
type ListQuotasRespQuotas struct {

	// GES资源配额列表。
	Resources *[]ListQuotasRespQuotasResources `json:"resources,omitempty"`
}

func (o ListQuotasRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRespQuotas struct{}"
	}

	return strings.Join([]string{"ListQuotasRespQuotas", string(data)}, " ")
}
