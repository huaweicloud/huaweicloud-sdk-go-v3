package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagCreate struct {

	// 键。 key最大长度为36个字符。 key不能为空字符串。 key前后空格会被丢弃。 key不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。 [key只能由中文，字母，数字，“-”，“_”组成。](tag:hws,hws_hk,fcs_vm,ctc)   [key只能由字母，数字，“_”，“-”组成。](tag:dt,ocb,tlf,sbc,g42,hcso_dt,tm) 创建时如果请求体中存在重复key则报错。 创建时，不允许重复key，如果数据库存在就覆盖。 删除时，允许重复key。 删除时，如果删除的标签不存在，默认处理成功,删除时不对标签字符集范围做校验。 删除时tags结构体不能缺失，key不能为空，或者空字符串。
	Key string `json:"key"`

	// 值。 添加标签时value值必选，删除标签时value值可选。 value最大长度为43个字符。 value可以为空字符串。 value前后的空格会被丢弃。 value不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。 [value只能由中文，字母，数字，“-”，“_”，“.”组成。](tag:hws,hws_hk,fcs_vm,ctc) [value只能由字母，数字，“_”，“-”组成。](tag:dt,ocb,tlf,sbc,g42,hcso_dt,tm)
	Value string `json:"value"`
}

func (o TagCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCreate struct{}"
	}

	return strings.Join([]string{"TagCreate", string(data)}, " ")
}
