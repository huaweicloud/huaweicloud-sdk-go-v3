package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建复制对请求体数据结构
type CreateReplicationRequestParams struct {

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`

	// 生产站点卷的ID。
	VolumeId string `json:"volume_id"`

	// 指定复制对的名称，最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name"`

	// 指定复制对的描述，最大支持长度为64个字节，不能包含左尖括号（<）或右尖括号（>）。
	Description *string `json:"description,omitempty"`

	// 专属分布式存储池ID。
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o CreateReplicationRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationRequestParams struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequestParams", string(data)}, " ")
}
