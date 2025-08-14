package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationTemplateDisplayDto 应用程序模板显示信息
type ApplicationTemplateDisplayDto struct {

	// 应用程序Id,以app-为前缀
	ApplicationId string `json:"application_id"`

	Display *DisplayDto `json:"display"`

	// 应用程序类型
	ApplicationType *string `json:"application_type,omitempty"`
}

func (o ApplicationTemplateDisplayDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationTemplateDisplayDto struct{}"
	}

	return strings.Join([]string{"ApplicationTemplateDisplayDto", string(data)}, " ")
}
