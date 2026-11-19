package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DelegatedAccountResponseInfo 已委托账号信息
type DelegatedAccountResponseInfo struct {

	// **参数解释** 父节点的ID **取值范围** 字符长度1-64位
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释** 父节点名称 **取值范围** 字符长度0-64位
	ParentName *string `json:"parent_name,omitempty"`

	// **参数解释** 账号ID **取值范围** 字符长度1-64位
	Id *string `json:"id,omitempty"`

	// **参数解释** 账号的统一资源名称 **取值范围** 字符长度0-2048位
	Urn *string `json:"urn,omitempty"`

	// **参数解释** 账号名称 **取值范围** 字符长度1-64位
	Name *string `json:"name,omitempty"`

	// **参数解释** 节点类型 **取值范围** - unit：组织 - account：账号 - project：项目
	OrgType *string `json:"org_type,omitempty"`
}

func (o DelegatedAccountResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DelegatedAccountResponseInfo struct{}"
	}

	return strings.Join([]string{"DelegatedAccountResponseInfo", string(data)}, " ")
}
