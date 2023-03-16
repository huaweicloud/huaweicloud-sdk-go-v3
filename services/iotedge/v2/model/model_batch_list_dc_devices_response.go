package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListDcDevicesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	Devices        *[]QueryDcDeviceRespDto `json:"devices,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchListDcDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcDevicesResponse struct{}"
	}

	return strings.Join([]string{"BatchListDcDevicesResponse", string(data)}, " ")
}
