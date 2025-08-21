package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRepoWebHookSubscriptionDto struct {

	// **参数解释：** webhook的url (需要base64编码)。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 秘钥。
	Token *string `json:"token,omitempty"`

	// **参数解释：** userid列表，提醒群中的指定成员（@某个成员），最长1000，每个最长100，“;”分隔符。
	MentionUsers *string `json:"mention_users,omitempty"`

	// **参数解释：** 手机号列表(需要base64编码)，提醒手机号对应的群成员（@某个成员），最长1000，每个最长30，“;”分隔符。
	MentionPhone *string `json:"mention_phone,omitempty"`
}

func (o UpdateRepoWebHookSubscriptionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoWebHookSubscriptionDto struct{}"
	}

	return strings.Join([]string{"UpdateRepoWebHookSubscriptionDto", string(data)}, " ")
}
