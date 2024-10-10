package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWhiteBlackIpRuleRequestBody 更新黑白名单规则请求体，其中请求体中必须包含name、white以及addr或者ip_group_id中一个
type UpdateWhiteBlackIpRuleRequestBody struct {

	// 黑白名单规则名称
	Name string `json:"name"`

	// 黑白名单ip地址，需要输入标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16
	Addr *string `json:"addr,omitempty"`

	// 黑白名单规则描述
	Description *string `json:"description,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White int32 `json:"white"`

	// 创建的Ip地址组id，该参数与addr参数使用一个即可；Ip地址组可在控制台中对象管理->地址组管理中添加。
	IpGroupId *string `json:"ip_group_id,omitempty"`

	// 生效模式，默认为permanent（立即生效）,创建自定义生效规则时请输入：customize
	TimeMode *string `json:"time_mode,omitempty"`

	// 规则生效开始时间，生效模式为自定义时，此字段才有效，请输入时间戳
	Start *int32 `json:"start,omitempty"`

	// 规则生效结束时间，生效模式为自定义时，此字段才有效，请输入时间戳
	Terminal *int32 `json:"terminal,omitempty"`
}

func (o UpdateWhiteBlackIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}
