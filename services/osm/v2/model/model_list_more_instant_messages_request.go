package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMoreInstantMessagesRequest Request Object
type ListMoreInstantMessagesRequest struct {

	// 工单id
	CaseId string `json:"case_id"`

	// 创建时间，时间戳
	CreateTime string `json:"create_time"`

	// 上一条消息的留言方式 0是客户留言 1是客服回留言
	Type int32 `json:"type"`

	// 华为云IAM组id，同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`

	// 查询数量
	Limit int32 `json:"limit"`

	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。
	XSite *int32 `json:"X-Site,omitempty"`

	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。
	XLanguage *string `json:"X-Language,omitempty"`

	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。
	XTimeZone *string `json:"X-Time-Zone,omitempty"`
}

func (o ListMoreInstantMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMoreInstantMessagesRequest struct{}"
	}

	return strings.Join([]string{"ListMoreInstantMessagesRequest", string(data)}, " ")
}
