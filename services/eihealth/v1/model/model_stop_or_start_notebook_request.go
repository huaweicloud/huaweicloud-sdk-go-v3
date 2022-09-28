package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopOrStartNotebookRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// notebook id
	NotebookId string `json:"notebook_id"`

	Body *NotebookActionReq `json:"body,omitempty"`
}

func (o StopOrStartNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopOrStartNotebookRequest struct{}"
	}

	return strings.Join([]string{"StopOrStartNotebookRequest", string(data)}, " ")
}
