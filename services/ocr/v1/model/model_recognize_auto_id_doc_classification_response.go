package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeAutoIdDocClassificationResponse Response Object
type RecognizeAutoIdDocClassificationResponse struct {
	Result *AutoIdDocClassificationResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeAutoIdDocClassificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeAutoIdDocClassificationResponse struct{}"
	}

	return strings.Join([]string{"RecognizeAutoIdDocClassificationResponse", string(data)}, " ")
}
