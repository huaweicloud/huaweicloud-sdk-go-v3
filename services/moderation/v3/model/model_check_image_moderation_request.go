package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckImageModerationRequest Request Object
type CheckImageModerationRequest struct {
	Body *ImageDetectionReq `json:"body,omitempty"`
}

func (o CheckImageModerationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckImageModerationRequest struct{}"
	}

	return strings.Join([]string{"CheckImageModerationRequest", string(data)}, " ")
}
