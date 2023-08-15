package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbInstanceRequest Request Object
type CreateDbInstanceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *OpenGaussInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateDbInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateDbInstanceRequest", string(data)}, " ")
}
