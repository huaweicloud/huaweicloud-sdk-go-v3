package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLabelPageRequest Request Object
type DeleteLabelPageRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 标签页面id
	LabelPageId string `json:"label_page_id"`
}

func (o DeleteLabelPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLabelPageRequest struct{}"
	}

	return strings.Join([]string{"DeleteLabelPageRequest", string(data)}, " ")
}
