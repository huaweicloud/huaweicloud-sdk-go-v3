package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteGaussMySqlDatabaseRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteGaussMySqlDatabaseRequestBody `json:"body,omitempty"`
}

func (o DeleteGaussMySqlDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlDatabaseRequest", string(data)}, " ")
}
