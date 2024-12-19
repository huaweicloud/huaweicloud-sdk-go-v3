package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseSchemaRequest Request Object
type DeleteDatabaseSchemaRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseSchemaRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseSchemaRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseSchemaRequest", string(data)}, " ")
}
