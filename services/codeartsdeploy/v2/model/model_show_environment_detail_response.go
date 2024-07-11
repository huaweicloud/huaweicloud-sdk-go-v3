package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvironmentDetailResponse Response Object
type ShowEnvironmentDetailResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *EnvironmentDetail `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowEnvironmentDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvironmentDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowEnvironmentDetailResponse", string(data)}, " ")
}
