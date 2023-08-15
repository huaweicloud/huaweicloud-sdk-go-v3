package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageMediaTaggingRequest Request Object
type RunImageMediaTaggingRequest struct {
	Body *ImageMediaTaggingReq `json:"body,omitempty"`
}

func (o RunImageMediaTaggingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMediaTaggingRequest struct{}"
	}

	return strings.Join([]string{"RunImageMediaTaggingRequest", string(data)}, " ")
}
