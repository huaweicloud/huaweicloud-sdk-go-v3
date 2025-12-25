package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceAgencyRequest Request Object
type CreateServiceAgencyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateServiceAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateServiceAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceAgencyRequest", string(data)}, " ")
}
