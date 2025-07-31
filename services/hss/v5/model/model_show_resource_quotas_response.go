package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceQuotasResponse Response Object
type ShowResourceQuotasResponse struct {

	// **参数解释**： 配额统计列表 **取值范围**： 不涉及
	DataList       *[]ResourceQuotasInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowResourceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceQuotasResponse", string(data)}, " ")
}
