package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOfflineRequest Request Object
type BatchOfflineRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *ApprovalBatchParam `json:"body,omitempty"`
}

func (o BatchOfflineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOfflineRequest struct{}"
	}

	return strings.Join([]string{"BatchOfflineRequest", string(data)}, " ")
}
