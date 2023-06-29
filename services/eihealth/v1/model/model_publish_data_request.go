package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishDataRequest Request Object
type PublishDataRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *PublishDataReq `json:"body,omitempty"`
}

func (o PublishDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishDataRequest struct{}"
	}

	return strings.Join([]string{"PublishDataRequest", string(data)}, " ")
}
