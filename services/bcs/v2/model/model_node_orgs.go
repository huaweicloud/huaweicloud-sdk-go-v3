package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点组织
type NodeOrgs struct {

	// 组织名称，IEF节点绑定模式下组织名与IEF节点名称保持一致。支持英文，数字，中文字符和中划线(-), 不能以中划线(-)开头，长度4-24个字符
	Name string `json:"name" xml:"name"`

	// 组织目标节点数, 1-2的正整数
	NodeCount int64 `json:"node_count" xml:"node_count"`

	// pvc名称，添加组织时需要提供pvc_name。CCE模式必填
	PvcName *string `json:"pvc_name,omitempty" xml:"pvc_name"`
}

func (o NodeOrgs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeOrgs struct{}"
	}

	return strings.Join([]string{"NodeOrgs", string(data)}, " ")
}
