package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGaussMySqlQuotasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *SetQuotasRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateGaussMySqlQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlQuotasRequest", string(data)}, " ")
}
