package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateSignatureRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *VerifyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ValidateSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSignatureRequest struct{}"
	}

	return strings.Join([]string{"ValidateSignatureRequest", string(data)}, " ")
}
