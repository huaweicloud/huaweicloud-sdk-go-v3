package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasResp **参数解释**： 资源配额。
type QuotasResp struct {

	// **参数解释**： 资源配额列表。
	Resources []ResourceQuotasResp `json:"resources"`
}

func (o QuotasResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasResp struct{}"
	}

	return strings.Join([]string{"QuotasResp", string(data)}, " ")
}
