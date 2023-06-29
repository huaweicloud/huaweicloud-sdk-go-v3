package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReplicationNameRequestParams 更新复制对名称请求体数据结构
type UpdateReplicationNameRequestParams struct {

	// 复制对的名称。最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name"`
}

func (o UpdateReplicationNameRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplicationNameRequestParams struct{}"
	}

	return strings.Join([]string{"UpdateReplicationNameRequestParams", string(data)}, " ")
}
