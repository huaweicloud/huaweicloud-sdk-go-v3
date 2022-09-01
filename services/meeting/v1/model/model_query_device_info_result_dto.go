package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备型号信息
type QueryDeviceInfoResultDto struct {

	// 终端型号
	Model *string `json:"model,omitempty" xml:"model"`

	// 设备终端产品尺寸
	DeviceSize *string `json:"deviceSize,omitempty" xml:"deviceSize"`

	// 终端设备购买渠道
	PurchaseChannel *string `json:"purchaseChannel,omitempty" xml:"purchaseChannel"`
}

func (o QueryDeviceInfoResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDeviceInfoResultDto struct{}"
	}

	return strings.Join([]string{"QueryDeviceInfoResultDto", string(data)}, " ")
}
