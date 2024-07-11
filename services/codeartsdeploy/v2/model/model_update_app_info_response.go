package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppInfoResponse Response Object
type UpdateAppInfoResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *AppBaseResponse `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateAppInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppInfoResponse", string(data)}, " ")
}
