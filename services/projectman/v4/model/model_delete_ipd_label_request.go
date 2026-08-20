package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpdLabelRequest Request Object
type DeleteIpdLabelRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	// 标签Id
	LabelId string `json:"label_id"`
}

func (o DeleteIpdLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpdLabelRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpdLabelRequest", string(data)}, " ")
}
