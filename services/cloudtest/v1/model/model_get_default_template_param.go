package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetDefaultTemplateParam struct {

	// 脑图模板名称
	Name *string `json:"name,omitempty"`
}

func (o GetDefaultTemplateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDefaultTemplateParam struct{}"
	}

	return strings.Join([]string{"GetDefaultTemplateParam", string(data)}, " ")
}
