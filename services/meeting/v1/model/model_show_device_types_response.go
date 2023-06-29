package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceTypesResponse Response Object
type ShowDeviceTypesResponse struct {

	// 查询到的终端类型列表
	Body           *[]QueryDeviceTypeResultDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowDeviceTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceTypesResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceTypesResponse", string(data)}, " ")
}
