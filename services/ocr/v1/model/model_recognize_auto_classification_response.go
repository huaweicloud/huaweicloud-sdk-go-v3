package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeAutoClassificationResponse struct {
	Result         *AutoClassificationResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RecognizeAutoClassificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeAutoClassificationResponse struct{}"
	}

	return strings.Join([]string{"RecognizeAutoClassificationResponse", string(data)}, " ")
}
