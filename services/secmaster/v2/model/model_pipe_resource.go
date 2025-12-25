package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipeResource struct {

	// 管道资源
	PipeResourceType *interface{} `json:"pipe_resource_type"`
}

func (o PipeResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeResource struct{}"
	}

	return strings.Join([]string{"PipeResource", string(data)}, " ")
}
