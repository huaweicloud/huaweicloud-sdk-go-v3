package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddUserRequestBody 用户信息。
type AddUserRequestBody struct {

	// 企业用户名称。
	Name string `json:"name"`

	// 后台自动识别是手机还是邮箱，若为手机号，则要求和国家码匹配。 > * 当前中国站点企业支持使用邮箱或手机号进行邀请，手机仅支持+86开头的手机号。 > * 当前国际站点企业仅支持使用邮箱进行邀请。
	Contact string `json:"contact"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 部门编码，若不携带则默认根部门。
	DeptCode *string `json:"deptCode,omitempty"`

	// 职位。
	Title *string `json:"title,omitempty"`

	// 通讯录排序等级，序号越低优先级越高。
	SortLevel *int32 `json:"sortLevel,omitempty"`

	// 备注。
	Desc *string `json:"desc,omitempty"`
}

func (o AddUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserRequestBody struct{}"
	}

	return strings.Join([]string{"AddUserRequestBody", string(data)}, " ")
}
