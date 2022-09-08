package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageTaggingRequest struct {
	Body *ImageTaggingReq `json:"body,omitempty"`
}

func (o RunImageTaggingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageTaggingRequest struct{}"
	}

	return strings.Join([]string{"RunImageTaggingRequest", string(data)}, " ")
}
