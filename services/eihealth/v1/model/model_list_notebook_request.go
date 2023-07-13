package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotebookRequest Request Object
type ListNotebookRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 读取条数
	Limit *int32 `json:"limit,omitempty"`

	// notebook名称
	Name *string `json:"name,omitempty"`

	// 读取偏移量
	Offset *int32 `json:"offset,omitempty"`

	// notebook状态
	Status *string `json:"status,omitempty"`
}

func (o ListNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotebookRequest struct{}"
	}

	return strings.Join([]string{"ListNotebookRequest", string(data)}, " ")
}
