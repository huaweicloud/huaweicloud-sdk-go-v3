package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineTagRequest Request Object
type DeletePipelineTagRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 标签ID
	TagId string `json:"tagId"`
}

func (o DeletePipelineTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineTagRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineTagRequest", string(data)}, " ")
}
