package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEvsQuotaResponse Response Object
type ShowEvsQuotaResponse struct {

	// 总配额
	Total *int32 `json:"total,omitempty"`

	// 已使用
	Usage *int32 `json:"usage,omitempty"`

	// 单位
	Unit           *string `json:"unit,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEvsQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvsQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowEvsQuotaResponse", string(data)}, " ")
}
