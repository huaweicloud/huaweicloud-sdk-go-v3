package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterAuthInfo 多云集群的账号权限
type MultiCloudClusterAuthInfo struct {

	// **参数解释**： API组 **取值范围**： 字符长度1-64位
	ApiGroups *string `json:"api_groups,omitempty"`

	// **参数解释**： 资源 **取值范围**： 字符长度1-16位
	Resources *string `json:"resources,omitempty"`

	// **参数解释**： 所属特性 **取值范围**： 字符长度1-16位
	Function *string `json:"function,omitempty"`

	// 是否有权访问
	IsAuthed *bool `json:"is_authed,omitempty"`

	// **参数解释**： 修复建议 **取值范围**： 字符长度1-128位
	Advice *string `json:"advice,omitempty"`
}

func (o MultiCloudClusterAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterAuthInfo struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterAuthInfo", string(data)}, " ")
}
