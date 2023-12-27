package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTextLanguageRequest Request Object
type CheckTextLanguageRequest struct {
	Body *LanguageCheckInfoReq `json:"body,omitempty"`
}

func (o CheckTextLanguageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTextLanguageRequest struct{}"
	}

	return strings.Join([]string{"CheckTextLanguageRequest", string(data)}, " ")
}
