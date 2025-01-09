package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateChangeImagesRequest Request Object
type EstimateChangeImagesRequest struct {
	Body *EstimateChangeImageRequestBody `json:"body,omitempty"`
}

func (o EstimateChangeImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateChangeImagesRequest struct{}"
	}

	return strings.Join([]string{"EstimateChangeImagesRequest", string(data)}, " ")
}
