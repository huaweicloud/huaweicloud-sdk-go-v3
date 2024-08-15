package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TicketInfoEnumData 单据枚举
type TicketInfoEnumData struct {

	// 字段KEY 标识哪个字段
	FiledKey string `json:"filed_key"`

	// 枚举KEY
	EnumKey string `json:"enum_key"`

	// 中文名称
	NameZh string `json:"name_zh"`

	// 英文名称
	NameEn string `json:"name_en"`
}

func (o TicketInfoEnumData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TicketInfoEnumData struct{}"
	}

	return strings.Join([]string{"TicketInfoEnumData", string(data)}, " ")
}
