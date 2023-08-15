package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasQuotas 配额列表对象。
type QuotasQuotas struct {

	// 资源列表对象。
	Resources *[]QuotasResource `json:"resources,omitempty"`
}

func (o QuotasQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasQuotas struct{}"
	}

	return strings.Join([]string{"QuotasQuotas", string(data)}, " ")
}
