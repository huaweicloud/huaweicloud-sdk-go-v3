package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
