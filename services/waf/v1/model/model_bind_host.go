package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 绑定的域名信息
type BindHost struct {

	// 域名ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 域名
	Hostname *string `json:"hostname,omitempty" xml:"hostname"`

	// 域名对应模式：cloud（云模式）/premium（独享模式）
	WafType *string `json:"waf_type,omitempty" xml:"waf_type"`

	// 仅独享模式涉及特殊域名模式
	Mode *string `json:"mode,omitempty" xml:"mode"`
}

func (o BindHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindHost struct{}"
	}

	return strings.Join([]string{"BindHost", string(data)}, " ")
}
