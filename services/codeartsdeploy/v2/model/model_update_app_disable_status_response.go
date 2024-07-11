package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppDisableStatusResponse Response Object
type UpdateAppDisableStatusResponse struct {

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAppDisableStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppDisableStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppDisableStatusResponse", string(data)}, " ")
}
