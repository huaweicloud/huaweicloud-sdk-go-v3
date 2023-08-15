package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageMediaTaggingDetRequest Request Object
type RunImageMediaTaggingDetRequest struct {
	Body *ImageMediaTaggingDetReq `json:"body,omitempty"`
}

func (o RunImageMediaTaggingDetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMediaTaggingDetRequest struct{}"
	}

	return strings.Join([]string{"RunImageMediaTaggingDetRequest", string(data)}, " ")
}
