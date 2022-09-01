package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// fault字段数据结构说明
type Fault struct {

	// 故障信息
	Message *string `json:"message,omitempty" xml:"message"`

	// 故障code
	Code *int32 `json:"code,omitempty" xml:"code"`

	// 故障详情
	Details *string `json:"details,omitempty" xml:"details"`

	// 故障时间。时间戳格式为ISO 8601：YYYY-MM-DDTHH:MM:SSZ，例如：2019-05-22T03:30:52Z
	Created *sdktime.SdkTime `json:"created,omitempty" xml:"created"`
}

func (o Fault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Fault struct{}"
	}

	return strings.Join([]string{"Fault", string(data)}, " ")
}
