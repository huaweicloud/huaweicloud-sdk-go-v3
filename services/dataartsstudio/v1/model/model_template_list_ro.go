package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateListRo struct {

	// 模板id列表
	Ids *[]int64 `json:"ids,omitempty"`
}

func (o TemplateListRo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateListRo struct{}"
	}

	return strings.Join([]string{"TemplateListRo", string(data)}, " ")
}
