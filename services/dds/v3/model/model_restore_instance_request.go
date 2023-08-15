package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInstanceRequest Request Object
type RestoreInstanceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
