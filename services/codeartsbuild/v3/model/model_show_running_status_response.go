package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRunningStatusResponse Response Object
type ShowRunningStatusResponse struct {
	Result *ShowRunningStatusResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRunningStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRunningStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowRunningStatusResponse", string(data)}, " ")
}
