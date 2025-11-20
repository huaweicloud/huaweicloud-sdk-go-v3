package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quotas struct {

	// - 参数解释：配额信息列表。 - 约束限制：不涉及。 - 取值范围：instance。 - 默认取值：不涉及。
	Resources []QuotaObject `json:"resources"`
}

func (o Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
