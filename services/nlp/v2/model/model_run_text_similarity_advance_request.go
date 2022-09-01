package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTextSimilarityAdvanceRequest struct {
	Body *TextSimilarityAdvanceRequest `json:"body,omitempty" xml:"body"`
}

func (o RunTextSimilarityAdvanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextSimilarityAdvanceRequest struct{}"
	}

	return strings.Join([]string{"RunTextSimilarityAdvanceRequest", string(data)}, " ")
}
