package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelfDefinedFieldVo 表自定义项。
type SelfDefinedFieldVo struct {

	// 自定义项中文名。
	FdNameCh *string `json:"fd_name_ch,omitempty"`

	// 自定义项英文名。
	FdNameEn *string `json:"fd_name_en,omitempty"`

	// 是否必填。
	NotNull *bool `json:"not_null,omitempty"`

	// 属性值。
	FdValue *string `json:"fd_value,omitempty"`
}

func (o SelfDefinedFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelfDefinedFieldVo struct{}"
	}

	return strings.Join([]string{"SelfDefinedFieldVo", string(data)}, " ")
}
