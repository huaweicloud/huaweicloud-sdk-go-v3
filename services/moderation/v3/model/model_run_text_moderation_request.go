package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTextModerationRequest Request Object
type RunTextModerationRequest struct {
	Body *TextDetectionReq `json:"body,omitempty"`
}

func (o RunTextModerationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextModerationRequest struct{}"
	}

	return strings.Join([]string{"RunTextModerationRequest", string(data)}, " ")
}
