package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseLabelInfo struct {

	// 标签标识
	LabelId *int32 `json:"label_id,omitempty" xml:"label_id"`

	// 标签名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 标签颜色
	Color *string `json:"color,omitempty" xml:"color"`
}

func (o CaseLabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseLabelInfo struct{}"
	}

	return strings.Join([]string{"CaseLabelInfo", string(data)}, " ")
}
