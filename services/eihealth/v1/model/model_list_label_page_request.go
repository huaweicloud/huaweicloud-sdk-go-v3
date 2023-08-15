package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelPageRequest Request Object
type ListLabelPageRequest struct {

	// 华为云项目ID，您可以从[获取项目ID](eihealth_33_0044.xml)中获取。
	EihealthProjectId string `json:"eihealth_project_id"`
}

func (o ListLabelPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelPageRequest struct{}"
	}

	return strings.Join([]string{"ListLabelPageRequest", string(data)}, " ")
}
