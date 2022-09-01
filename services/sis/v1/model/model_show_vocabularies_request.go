package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVocabulariesRequest struct {
	Body *ShowVocabulariesParams `json:"body,omitempty" xml:"body"`
}

func (o ShowVocabulariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabulariesRequest struct{}"
	}

	return strings.Join([]string{"ShowVocabulariesRequest", string(data)}, " ")
}
