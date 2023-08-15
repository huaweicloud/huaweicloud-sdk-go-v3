package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCorpAdminResultDto struct {

	// 用户UUID。
	Id *string `json:"id,omitempty"`

	// 用户帐号（华为云会议帐号）。
	Account *string `json:"account,omitempty"`

	// 用户名称。
	Name *string `json:"name,omitempty"`

	// 管理员类型。 * 0：默认管理员 * 1：普通管理员
	AdminType *int32 `json:"adminType,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 手机号。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	Dept *DeptBasicDto `json:"dept,omitempty"`
}

func (o QueryCorpAdminResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpAdminResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpAdminResultDto", string(data)}, " ")
}
