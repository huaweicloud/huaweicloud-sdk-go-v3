package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobConfigResponse Response Object
type ShowJobConfigResponse struct {
	Result *CreateBuildJobRequestBody `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowJobConfigResponse", string(data)}, " ")
}
