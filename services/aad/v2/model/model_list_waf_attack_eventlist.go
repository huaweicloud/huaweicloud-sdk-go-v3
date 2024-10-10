package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListWafAttackEventlist struct {

	// id
	Id *string `json:"id,omitempty"`

	// 攻击目标域名
	Domain *string `json:"domain,omitempty"`

	// 攻击时间
	Time *int32 `json:"time,omitempty"`

	// 攻击源IP
	Sip *string `json:"sip,omitempty"`

	// 防御动作
	Action *string `json:"action,omitempty"`

	// 攻击url
	Url *string `json:"url,omitempty"`

	// 攻击类型
	Type *string `json:"type,omitempty"`

	Backend *Backend `json:"backend,omitempty"`
}

func (o ListWafAttackEventlist) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafAttackEventlist struct{}"
	}

	return strings.Join([]string{"ListWafAttackEventlist", string(data)}, " ")
}
