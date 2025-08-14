package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppExtendedInfo 包含自定义应用业务扩展信息的对象
type AppExtendedInfo struct {

	// 扩展信息的键值对映射
	ExtendedInfo map[string]string `json:"extended_info,omitempty"`
}

func (o AppExtendedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppExtendedInfo struct{}"
	}

	return strings.Join([]string{"AppExtendedInfo", string(data)}, " ")
}
