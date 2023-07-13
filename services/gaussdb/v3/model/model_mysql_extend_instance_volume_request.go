package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlExtendInstanceVolumeRequest 扩容信息
type MysqlExtendInstanceVolumeRequest struct {

	// 扩容后的容量。包周期实例初始最小磁盘规格为10G，实例所选容量大小必须为10的整数倍，且大于实际使用容量，最大为128000GB.  取值范围： 扩容时必须大于等于20G； 缩容时必须大于等于10G。
	Size int32 `json:"size"`

	// 表示是否自动从客户的账户中支付。  - true，为自动支付，默认该方式。 - false，为手动支付。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o MysqlExtendInstanceVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlExtendInstanceVolumeRequest struct{}"
	}

	return strings.Join([]string{"MysqlExtendInstanceVolumeRequest", string(data)}, " ")
}
