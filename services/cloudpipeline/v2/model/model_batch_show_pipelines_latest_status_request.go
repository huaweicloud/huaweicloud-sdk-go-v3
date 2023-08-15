package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPipelinesLatestStatusRequest Request Object
type BatchShowPipelinesLatestStatusRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchShowPipelinesLatestStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPipelinesLatestStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchShowPipelinesLatestStatusRequest", string(data)}, " ")
}
