package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTemplateReqNew struct {

	// 模板名称。
	Name string `json:"name"`
}

func (o DeleteTemplateReqNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateReqNew struct{}"
	}

	return strings.Join([]string{"DeleteTemplateReqNew", string(data)}, " ")
}
