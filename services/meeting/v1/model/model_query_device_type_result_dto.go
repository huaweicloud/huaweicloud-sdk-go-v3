package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDeviceTypeResultDto 终端类型。
type QueryDeviceTypeResultDto struct {

	// 终端类型，区分自研和第三方终端。 * TE：华为自研硬终端 * 3rd：第三方硬终端
	Type *string `json:"type,omitempty"`

	// 终端型号，枚举类型。 * TE10 * TE20 * TE30 * TE40 * TE50 * TE60 * HUAWEI Box 300 * HUAWEI Box 500 * HUAWEI Box 600 * HUAWEI Box 700 * HUAWEI Box 900 * DP300 * HUAWEI Box 200 * HUAWEI Box 300 * HUAWEI Box 500 * HUAWEI Board * polycomcisco
	Model *string `json:"model,omitempty"`

	// 是否支持激活码。
	EnableActiveCode *bool `json:"enableActiveCode,omitempty"`

	// 屏幕分辨率。1080P、720P等。
	Resolution *string `json:"resolution,omitempty"`

	// 是否支持投影码。
	SupportProjectionCode *bool `json:"supportProjectionCode,omitempty"`

	// 是否支持SVC。
	SupportSVC *bool `json:"supportSVC,omitempty"`
}

func (o QueryDeviceTypeResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeviceTypeResultDto struct{}"
	}

	return strings.Join([]string{"QueryDeviceTypeResultDto", string(data)}, " ")
}
