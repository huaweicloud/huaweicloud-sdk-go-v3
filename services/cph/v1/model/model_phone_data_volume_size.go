package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PhoneDataVolumeSize struct {

	// 云手机id
	PhoneId string `json:"phone_id"`

	// 扩容后的手机数据盘大小，单位为GiB。扩容的大小必须大于等于原有容量且小于最大容量。 最大容量：32768GiB
	NewSize int32 `json:"new_size"`
}

func (o PhoneDataVolumeSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneDataVolumeSize struct{}"
	}

	return strings.Join([]string{"PhoneDataVolumeSize", string(data)}, " ")
}
