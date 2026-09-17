package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasResponse Response Object
type ShowQuotasResponse struct {

	// **参数解释：** 资源配额列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Quotas         *[]QuotaResource `json:"quotas,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
