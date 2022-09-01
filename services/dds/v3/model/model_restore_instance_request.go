package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreInstanceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *RestoreInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o RestoreInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
