package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssuesTicketInfoEnumData 单据枚举。
type IssuesTicketInfoEnumData struct {

	// 是否已删除。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 匹配的枚举类型。
	MatchType *string `json:"match_type,omitempty"`

	// 当前工单ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 工单单号。
	RealTicketId *string `json:"real_ticket_id,omitempty"`

	// 中文名称。
	NameZh *string `json:"name_zh,omitempty"`

	// 英文名称。
	NameEn *string `json:"name_en,omitempty"`

	// 枚举值对应的唯一id，当match_type为reference__base_config.User时，biz_id的值为操作用户唯一Id。
	BizId *string `json:"biz_id,omitempty"`

	// 当前枚举值对应的类型。
	PropId *string `json:"prop_id,omitempty"`

	// 后台不同应用对应的模型id。
	ModelId *string `json:"model_id,omitempty"`
}

func (o IssuesTicketInfoEnumData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssuesTicketInfoEnumData struct{}"
	}

	return strings.Join([]string{"IssuesTicketInfoEnumData", string(data)}, " ")
}
