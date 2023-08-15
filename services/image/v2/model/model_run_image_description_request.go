package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageDescriptionRequest Request Object
type RunImageDescriptionRequest struct {
	Body *ImageDescriptionReq `json:"body,omitempty"`
}

func (o RunImageDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageDescriptionRequest struct{}"
	}

	return strings.Join([]string{"RunImageDescriptionRequest", string(data)}, " ")
}
