package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationV4Request Request Object
type DeleteApplicationV4Request struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 是否删除代码仓
	IsDeleteRepository *bool `json:"is_delete_repository,omitempty"`

	// 删除流水线ID,多流水线逗号隔开
	PipelineIds *string `json:"pipeline_ids,omitempty"`
}

func (o DeleteApplicationV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationV4Request struct{}"
	}

	return strings.Join([]string{"DeleteApplicationV4Request", string(data)}, " ")
}
