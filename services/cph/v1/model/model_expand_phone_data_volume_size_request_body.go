package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandPhoneDataVolumeSizeRequestBody 扩容云手机数据盘大小请求体。
type ExpandPhoneDataVolumeSizeRequestBody struct {

	// 云手机列表。
	Phones []PhoneDataVolumeSize `json:"phones"`
}

func (o ExpandPhoneDataVolumeSizeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandPhoneDataVolumeSizeRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandPhoneDataVolumeSizeRequestBody", string(data)}, " ")
}
