package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceTemplateUpdateDetailTags 设备模板标签，key-value键值对形式。
type DeviceTemplateUpdateDetailTags struct {

	// 标签key值，长度取值范围为1~36， 仅允许大小写英文字母、数字、下划线、中划线
	Key *string `json:"key,omitempty"`

	// 标签value值，长度取值范围为0~43， 仅允许大小写英文字母、数字、下划线、中划线
	Value *string `json:"value,omitempty"`
}

func (o DeviceTemplateUpdateDetailTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceTemplateUpdateDetailTags struct{}"
	}

	return strings.Join([]string{"DeviceTemplateUpdateDetailTags", string(data)}, " ")
}
