package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCasesRequest struct {
	// 工单id

	CaseId string `json:"case_id"`
	// 具体操作 cancel,close,press,delete

	ActionId string `json:"action_id"`
	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。

	XSite *int32 `json:"X-Site,omitempty"`
	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。

	XLanguage *string `json:"X-Language,omitempty"`
	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。

	XTimeZone *string `json:"X-Time-Zone,omitempty"`

	Body *WorkOrderOperateV2Req `json:"body,omitempty"`
}

func (o UpdateCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCasesRequest struct{}"
	}

	return strings.Join([]string{"UpdateCasesRequest", string(data)}, " ")
}
