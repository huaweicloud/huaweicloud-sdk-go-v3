package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MenuItem struct {

	// 子菜单名称。  1. 一级菜单名长度和菜单数量有关，具体约束为：     - 当菜单数量为1个时，菜单名长度范围在1-24个字符。    - 当菜单数量为2个时，菜单名长度范围在1-12个字符。    - 当菜单数量为3个时，菜单名长度范围在1-8个字符。  2. 二级菜单名长度范围恒为1-16个字符。  > 以上字符区分中英文，一个中文占2个字符，字母和数字占1个字符，且同时生效的一组菜单内名称不能重复。
	Name string `json:"name"`

	// 菜单动作类型。 - OPEN_SUBMENU：打开子菜单 - OPEN_URL：打开URL - CALLING：拨打电话 - OPEN_APP：打开APP - OPEN_QUICK：打开快应用
	ActionType string `json:"action_type"`

	// 对应值类型。对应不同action_type值，content含义如下： - action_type=OPEN_SUBMENU：不填 - action_type=OPEN_URL：参数数值为跳转URL - action_type=CALLING：参数数值为电话号码 - action_type=OPEN_APP：参数数值为APP的跳转deeplink - action_type=OPEN_QUICK：参数数值为快应用跳转的deeplink
	Content *string `json:"content,omitempty"`

	ExtMsg *ExtMsg `json:"ext_msg,omitempty"`

	// 子菜单配置项。  > 仅当action_type=OPEN_SUBMENU时生效，且该项内不允许再配置子菜单。
	SubMenuItems *[]MenuItem `json:"sub_menu_items,omitempty"`
}

func (o MenuItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MenuItem struct{}"
	}

	return strings.Join([]string{"MenuItem", string(data)}, " ")
}
