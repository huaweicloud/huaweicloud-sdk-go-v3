package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoWebHookSubscriptionDto struct {

	// **参数解释：** webhook的url。
	Url *string `json:"url,omitempty"`

	// **参数解释：** userid列表，提醒群中的指定成员（@某个成员），最长1000，每个最长100，“;”分隔符。
	MentionUsers *string `json:"mention_users,omitempty"`

	// **参数解释：** 手机号列表，提醒手机号对应的群成员（@某个成员），最长1000，每个最长30，“;”分隔符。
	MentionPhone *string `json:"mention_phone,omitempty"`

	// **参数解释：** 是否设置了token。
	HasToken *bool `json:"has_token,omitempty"`
}

func (o RepoWebHookSubscriptionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoWebHookSubscriptionDto struct{}"
	}

	return strings.Join([]string{"RepoWebHookSubscriptionDto", string(data)}, " ")
}
