package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPublishRequest Request Object
type BatchPublishRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *ApprovalBatchParam `json:"body,omitempty"`
}

func (o BatchPublishRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublishRequest struct{}"
	}

	return strings.Join([]string{"BatchPublishRequest", string(data)}, " ")
}
