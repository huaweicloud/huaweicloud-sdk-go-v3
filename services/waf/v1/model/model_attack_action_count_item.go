package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttackActionCountItem struct {

	// **参数解释：** 防护动作 **取值范围：** - block: 拦截，表示超过“限速频率”将直接拦截。 - log：仅记录，表示超过“限速频率”将只记录不拦截。 - captcha：表示超过“限速频率”后弹出验证码，进行人机验证，完成验证后，请求将不受访问限制。人机验证目前支持英文。 - dynamic_block：上一个限速周期内，请求频率超过“限速频率”将被拦截，那么在下一个限速周期内，请求频率超过“放行频率”将被拦截。 - advanced_captcha：高阶人机验证，表示超过“限速频率”后弹出验证码，进行人机验证。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 数量 **取值范围：** 不涉及
	Num *int64 `json:"num,omitempty"`
}

func (o AttackActionCountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackActionCountItem struct{}"
	}

	return strings.Join([]string{"AttackActionCountItem", string(data)}, " ")
}
