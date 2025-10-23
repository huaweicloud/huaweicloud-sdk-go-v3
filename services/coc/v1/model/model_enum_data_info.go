package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnumDataInfo 基础数据
type EnumDataInfo struct {

	// 是否删除
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 匹配类型
	MatchType *string `json:"match_type,omitempty"`

	// 所属单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 所属真实单号。
	RealTicketId *string `json:"real_ticket_id,omitempty"`

	// 中文名
	NameZh *string `json:"name_zh,omitempty"`

	// 英文名
	NameEn *string `json:"name_en,omitempty"`

	// 工号
	UserName *string `json:"user_name,omitempty"`

	// 唯一id
	BizId *string `json:"biz_id,omitempty"`

	// 字段id
	PropId *string `json:"prop_id,omitempty"`

	// 模型id
	ModelId *string `json:"model_id,omitempty"`
}

func (o EnumDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnumDataInfo struct{}"
	}

	return strings.Join([]string{"EnumDataInfo", string(data)}, " ")
}
