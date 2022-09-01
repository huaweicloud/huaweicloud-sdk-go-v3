package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新保护实例名称请求数据结构
type UpdateProtectedInstanceNameRequestParams struct {

	// 保护实例的名称。最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name" xml:"name"`
}

func (o UpdateProtectedInstanceNameRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameRequestParams struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameRequestParams", string(data)}, " ")
}
