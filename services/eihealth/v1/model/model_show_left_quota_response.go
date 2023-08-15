package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLeftQuotaResponse Response Object
type ShowLeftQuotaResponse struct {

	// 剩余服务器数
	Quantity *int32 `json:"quantity,omitempty"`

	// 剩余CPU
	Cpu *int32 `json:"cpu,omitempty"`

	// 剩余内存
	Ram            *int32 `json:"ram,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowLeftQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLeftQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowLeftQuotaResponse", string(data)}, " ")
}
