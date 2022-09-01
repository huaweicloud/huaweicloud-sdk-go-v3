package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageDescriptionRequest struct {
	Body *ImageDescriptionReq `json:"body,omitempty" xml:"body"`
}

func (o RunImageDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageDescriptionRequest struct{}"
	}

	return strings.Join([]string{"RunImageDescriptionRequest", string(data)}, " ")
}