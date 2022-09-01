package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchDevicesResponse struct {

	// 页面起始页，从0开始
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 总数量。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 终端信息列表。
	Data           *[]QueryDeviceResultDto `json:"data,omitempty" xml:"data"`
	HttpStatusCode int                     `json:"-"`
}

func (o SearchDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDevicesResponse struct{}"
	}

	return strings.Join([]string{"SearchDevicesResponse", string(data)}, " ")
}
