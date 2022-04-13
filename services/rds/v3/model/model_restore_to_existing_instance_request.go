package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreToExistingInstanceRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreToExistingInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreToExistingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreToExistingInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreToExistingInstanceRequest", string(data)}, " ")
}
