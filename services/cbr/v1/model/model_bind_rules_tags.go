package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自动绑定规则标签
type BindRulesTags struct {

	// key不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\\\”,“,”,“|”,“/”。 [key只能由中文，字母，数字，“-”，“_”组成。](tag:hws,hws_hk,fcs_vm,ctc)   [key只能由字母，数字，“_”，“-”组成。](tag:dt,ocb,tlf,sbc,g42,hcso_dt)
	Key string `json:"key"`

	// value不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。  [value只能由中文，字母，数字，“-”，“_”，“.”组成。](tag:hws,hws_hk,fcs_vm,ctc) [value只能由字母，数字，“_”，“-”组成。](tag:dt,ocb,tlf,sbc,g42,hcso_dt)
	Value *string `json:"value,omitempty"`
}

func (o BindRulesTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindRulesTags struct{}"
	}

	return strings.Join([]string{"BindRulesTags", string(data)}, " ")
}
