package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostDetailResponse Response Object
type ShowHostDetailResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *HostInfoDetail `json:"result,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHostDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHostDetailResponse", string(data)}, " ")
}
