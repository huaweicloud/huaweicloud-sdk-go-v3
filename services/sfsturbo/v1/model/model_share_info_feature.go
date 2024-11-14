package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShareInfoFeature 文件系统的特性支持情况。
type ShareInfoFeature struct {

	// 文件系统是否支持该特性
	IsSupport *bool `json:"is_support,omitempty"`

	// 文件系统是否支持该特性的详细信息
	Message *string `json:"message,omitempty"`

	// 文件系统是否支持该特性的详细信息
	MsgCode *string `json:"msg_code,omitempty"`
}

func (o ShareInfoFeature) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareInfoFeature struct{}"
	}

	return strings.Join([]string{"ShareInfoFeature", string(data)}, " ")
}
