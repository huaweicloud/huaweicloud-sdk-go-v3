package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTableFieldValueVo 码表属性值。
type CodeTableFieldValueVo struct {

	// 码表字段ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 所属码表属性ID，填写String类型替代Long类型。
	FdId *string `json:"fd_id,omitempty"`

	// 码表属性值。
	FdValue *string `json:"fd_value,omitempty"`

	// 序号。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o CodeTableFieldValueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTableFieldValueVo struct{}"
	}

	return strings.Join([]string{"CodeTableFieldValueVo", string(data)}, " ")
}
