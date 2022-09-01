package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeviceGroupTreeResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 设备分组信息
	Items          *[]GroupTreeResponse `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowDeviceGroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupTreeResponse", string(data)}, " ")
}
