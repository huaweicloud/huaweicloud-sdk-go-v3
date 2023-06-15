package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGaussMySqlConfigurationRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateConfigurationRequestBody `json:"body,omitempty"`
}

func (o CreateGaussMySqlConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlConfigurationRequest", string(data)}, " ")
}
