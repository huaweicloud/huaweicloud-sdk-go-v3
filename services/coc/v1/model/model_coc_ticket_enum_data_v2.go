package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocTicketEnumDataV2 TicketEnumData
type CocTicketEnumDataV2 struct {

	// 匹配属性字段
	PropId *string `json:"prop_id,omitempty"`

	// 字段值
	BizId *string `json:"biz_id,omitempty"`

	// 中文名称
	NameZh *string `json:"name_zh,omitempty"`

	// 英文名称
	NameEn *string `json:"name_en,omitempty"`
}

func (o CocTicketEnumDataV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTicketEnumDataV2 struct{}"
	}

	return strings.Join([]string{"CocTicketEnumDataV2", string(data)}, " ")
}
