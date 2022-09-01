package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertyRequest struct {

	// 属性名称，必须是模型中已存在的
	Name string `json:"name" xml:"name"`

	// 值，只有static型属性可以填写
	Value *interface{} `json:"value,omitempty" xml:"value"`

	// 设备ID，只有measurement型属性可以填写
	DeviceId *string `json:"device_id,omitempty" xml:"device_id"`
}

func (o PropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyRequest struct{}"
	}

	return strings.Join([]string{"PropertyRequest", string(data)}, " ")
}
