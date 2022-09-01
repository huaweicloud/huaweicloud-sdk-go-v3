package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建黑白名单规则body
type CreateWhiteBlackIpRuleRequestBody struct {

	// 规则名称只能由字母、数字、-、_和.组成，长度不能超过64个字符
	Name string `json:"name" xml:"name"`

	// 黑白名单ip地址，需要输入标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16
	Addr *string `json:"addr,omitempty" xml:"addr"`

	// 黑白名单规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 防护动作：  - 0 拦截  - 1 放行   - 2 仅记录
	White int32 `json:"white" xml:"white"`

	// 创建的Ip地址组id，该参数与addr参数只能使用一个；Ip地址组可在控制台中对象管理->地址组管理中添加。
	IpGroupId *string `json:"ip_group_id,omitempty" xml:"ip_group_id"`
}

func (o CreateWhiteBlackIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}
