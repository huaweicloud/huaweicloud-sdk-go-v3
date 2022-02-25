package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端类型
type QueryDeviceTypeResultDto struct {
	// 终端类型，区分自研和第三方终端

	Type *string `json:"type,omitempty"`
	// 终端型号，枚举类型。当前支持TE系列硬件终端，具体的终端类型可以通过获取所有终端类型接口查询。

	Model *string `json:"model,omitempty"`
	// 是否支持激活码

	EnableActiveCode *bool `json:"enableActiveCode,omitempty"`
	// 屏幕分辨率。1080P、720P等。

	Resolution *string `json:"resolution,omitempty"`
	// 是否支持投影码

	SupportProjectionCode *bool `json:"supportProjectionCode,omitempty"`
	// 是否支持SVC

	SupportSVC *bool `json:"supportSVC,omitempty"`
}

func (o QueryDeviceTypeResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeviceTypeResultDto struct{}"
	}

	return strings.Join([]string{"QueryDeviceTypeResultDto", string(data)}, " ")
}
