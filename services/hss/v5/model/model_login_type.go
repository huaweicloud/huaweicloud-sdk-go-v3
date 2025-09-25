package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginType **参数解释**： 登录类型 **约束限制**: 不涉及 **取值范围**: - mysql：mysql服务。 - rdp：rdp服务。 - ssh：ssh服务。 - vsftp：vsftp服务。  **默认取值**: 不涉及
type LoginType struct {
}

func (o LoginType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginType struct{}"
	}

	return strings.Join([]string{"LoginType", string(data)}, " ")
}
