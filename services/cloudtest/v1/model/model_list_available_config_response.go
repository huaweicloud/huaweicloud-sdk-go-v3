package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableConfigResponse Response Object
type ListAvailableConfigResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorAvailableConfig `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	Result *AvailableConfig `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAvailableConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableConfigResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableConfigResponse", string(data)}, " ")
}
