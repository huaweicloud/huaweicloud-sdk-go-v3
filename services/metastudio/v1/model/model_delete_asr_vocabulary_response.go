package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAsrVocabularyResponse Response Object
type DeleteAsrVocabularyResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAsrVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAsrVocabularyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAsrVocabularyResponse", string(data)}, " ")
}
