package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseInstanceRequest Request Object
type CreateDatabaseInstanceRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CreateDatabaseInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateDatabaseInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseInstanceRequest", string(data)}, " ")
}
