package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTextTranslationRequest Request Object
type RunTextTranslationRequest struct {
	Body *TextTranslationReq `json:"body,omitempty"`
}

func (o RunTextTranslationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextTranslationRequest struct{}"
	}

	return strings.Join([]string{"RunTextTranslationRequest", string(data)}, " ")
}
