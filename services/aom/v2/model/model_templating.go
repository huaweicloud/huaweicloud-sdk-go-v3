package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Templating struct {

	// 变量列表。
	List *[]TemplateInfo `json:"list,omitempty"`
}

func (o Templating) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Templating struct{}"
	}

	return strings.Join([]string{"Templating", string(data)}, " ")
}
