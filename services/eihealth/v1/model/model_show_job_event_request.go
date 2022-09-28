package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobEventRequest struct {

	// Locale语言类型，zh_cn返回中文，en_us返回英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowJobEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobEventRequest struct{}"
	}

	return strings.Join([]string{"ShowJobEventRequest", string(data)}, " ")
}
