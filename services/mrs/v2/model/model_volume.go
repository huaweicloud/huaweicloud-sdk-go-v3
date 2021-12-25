package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Volume struct {
	// 磁盘类型。 磁盘类型枚举值： - SATA：普通IO磁盘类型。 - SAS：高IO磁盘类型。 - SSD：超高IO磁盘类型。 - GPSSD：通用型SSD磁盘类型

	Type string `json:"type"`
	// 数据盘大小，容量单位为GB，输入大小范围为[10,32768]。

	Size int32 `json:"size"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
