package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeAutoClassificationRequest struct {
	Body *AutoClassificationRequestBody `json:"body,omitempty"`
}

func (o RecognizeAutoClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeAutoClassificationRequest struct{}"
	}

	return strings.Join([]string{"RecognizeAutoClassificationRequest", string(data)}, " ")
}
