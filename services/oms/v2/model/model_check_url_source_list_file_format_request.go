package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckUrlSourceListFileFormatRequest Request Object
type CheckUrlSourceListFileFormatRequest struct {
	Body *CheckUrlSourceListFileFormatReq `json:"body,omitempty"`
}

func (o CheckUrlSourceListFileFormatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUrlSourceListFileFormatRequest struct{}"
	}

	return strings.Join([]string{"CheckUrlSourceListFileFormatRequest", string(data)}, " ")
}
