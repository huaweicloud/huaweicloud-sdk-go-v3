package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppDetailByIdResponse Response Object
type ShowAppDetailByIdResponse struct {
	Result *AppDetailInfo `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppDetailByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppDetailByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAppDetailByIdResponse", string(data)}, " ")
}
