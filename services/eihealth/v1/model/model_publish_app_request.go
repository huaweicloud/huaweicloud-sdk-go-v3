package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppRequest Request Object
type PublishAppRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 应用id
	AppId string `json:"app_id"`

	Body *PublishAppReq `json:"body,omitempty"`
}

func (o PublishAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppRequest struct{}"
	}

	return strings.Join([]string{"PublishAppRequest", string(data)}, " ")
}
