package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewSessionForKillProcessTaskNewRequest Request Object
type PreviewSessionForKillProcessTaskNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *PreviewSessionForKillProcessTaskNewRequestBody `json:"body,omitempty"`
}

func (o PreviewSessionForKillProcessTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewSessionForKillProcessTaskNewRequest struct{}"
	}

	return strings.Join([]string{"PreviewSessionForKillProcessTaskNewRequest", string(data)}, " ")
}
