package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionsResponse Response Object
type ShowPartitionsResponse struct {

	// 请求结果，true为成功，false为失败
	IsSuccess *bool `json:"is_success,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	Partitions     *Partitions `json:"partitions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionsResponse struct{}"
	}

	return strings.Join([]string{"ShowPartitionsResponse", string(data)}, " ")
}
