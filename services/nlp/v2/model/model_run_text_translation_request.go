package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunTextTranslationRequest struct {
	Body *TextTranslationReq `json:"body,omitempty" xml:"body"`
}

func (o RunTextTranslationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextTranslationRequest struct{}"
	}

	return strings.Join([]string{"RunTextTranslationRequest", string(data)}, " ")
}
