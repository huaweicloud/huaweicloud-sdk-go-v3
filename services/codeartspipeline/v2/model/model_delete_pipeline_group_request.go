package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipelineGroupRequest Request Object
type DeletePipelineGroupRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 分组ID
	Id string `json:"id"`
}

func (o DeletePipelineGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipelineGroupRequest struct{}"
	}

	return strings.Join([]string{"DeletePipelineGroupRequest", string(data)}, " ")
}
