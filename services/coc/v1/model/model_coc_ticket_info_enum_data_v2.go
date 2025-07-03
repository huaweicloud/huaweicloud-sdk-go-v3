package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocTicketInfoEnumDataV2 单据枚举
type CocTicketInfoEnumDataV2 struct {

	// 字段KEY 标识哪个字段
	PropId *string `json:"prop_id,omitempty"`

	// 枚举KEY
	BizId *string `json:"biz_id,omitempty"`

	// 中文名称
	NameZh string `json:"name_zh"`

	// 英文名称
	NameEn string `json:"name_en"`
}

func (o CocTicketInfoEnumDataV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTicketInfoEnumDataV2 struct{}"
	}

	return strings.Join([]string{"CocTicketInfoEnumDataV2", string(data)}, " ")
}
