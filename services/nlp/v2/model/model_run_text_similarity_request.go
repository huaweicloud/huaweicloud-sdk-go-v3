package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTextSimilarityRequest Request Object
type RunTextSimilarityRequest struct {
	Body *TextSimilarityRequest `json:"body,omitempty"`
}

func (o RunTextSimilarityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextSimilarityRequest struct{}"
	}

	return strings.Join([]string{"RunTextSimilarityRequest", string(data)}, " ")
}
