package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceQuotasResponse Response Object
type ShowResourceQuotasResponse struct {

	// 配额统计列表
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
