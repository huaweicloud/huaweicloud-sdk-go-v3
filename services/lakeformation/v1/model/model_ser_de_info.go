package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SerDeInfo 序列化/反序列化的信息
type SerDeInfo struct {

	// 名字信息
	Name *string `json:"name,omitempty"`

	// 实现序列化/反序列化的类
	SerializationLibrary *string `json:"serialization_library,omitempty"`

	// 参数数组。 key最小值1，最大值255 value最大值4000
	Parameters map[string]string `json:"parameters,omitempty"`
}

func (o SerDeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SerDeInfo struct{}"
	}

	return strings.Join([]string{"SerDeInfo", string(data)}, " ")
}
