package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunFileTranslationRequest struct {
	Body *FileTranslationReq `json:"body,omitempty" xml:"body"`
}

func (o RunFileTranslationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFileTranslationRequest struct{}"
	}

	return strings.Join([]string{"RunFileTranslationRequest", string(data)}, " ")
}
