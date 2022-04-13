package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunFileTranslationRequest struct {
	Body *FileTranslationReq `json:"body,omitempty"`
}

func (o RunFileTranslationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFileTranslationRequest struct{}"
	}

	return strings.Join([]string{"RunFileTranslationRequest", string(data)}, " ")
}
