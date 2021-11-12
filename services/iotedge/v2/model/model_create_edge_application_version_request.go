package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEdgeApplicationVersionRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`

	Body *CreateEdgeApplicationVersionDto `json:"body,omitempty"`
}

func (o CreateEdgeApplicationVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionRequest", string(data)}, " ")
}
