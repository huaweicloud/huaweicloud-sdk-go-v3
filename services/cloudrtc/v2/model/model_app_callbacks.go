package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// app回调配置
type AppCallbacks struct {
	PushCallback *AppCallbackUrl `json:"push_callback,omitempty" xml:"push_callback"`

	RecordCallback *AppCallbackUrl `json:"record_callback,omitempty" xml:"record_callback"`
}

func (o AppCallbacks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCallbacks struct{}"
	}

	return strings.Join([]string{"AppCallbacks", string(data)}, " ")
}
