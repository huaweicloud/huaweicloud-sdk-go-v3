package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseQueryLabel struct {

	// 标签标识
	LabelId *string `json:"labelId,omitempty"`

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 标签颜色
	Color *string `json:"color,omitempty"`
}

func (o CaseQueryLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseQueryLabel struct{}"
	}

	return strings.Join([]string{"CaseQueryLabel", string(data)}, " ")
}
