package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmVersionToGalleryRequest Request Object
type CreateAlgorithmVersionToGalleryRequest struct {
	Body *CreateAlgorithmVersionToGalleryBody `json:"body,omitempty"`
}

func (o CreateAlgorithmVersionToGalleryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmVersionToGalleryRequest struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmVersionToGalleryRequest", string(data)}, " ")
}
