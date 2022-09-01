package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreExistInstanceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *RestoreExistingInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RestoreExistInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreExistInstanceRequest", string(data)}, " ")
}
