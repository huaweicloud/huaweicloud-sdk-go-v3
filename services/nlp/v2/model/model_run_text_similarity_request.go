package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTextSimilarityRequest struct {
	Body *TextSimilarityRequest `json:"body,omitempty" xml:"body"`
}

func (o RunTextSimilarityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextSimilarityRequest struct{}"
	}

	return strings.Join([]string{"RunTextSimilarityRequest", string(data)}, " ")
}
