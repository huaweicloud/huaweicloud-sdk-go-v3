package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlDatabaseUserCommentRequest Request Object
type UpdateGaussMySqlDatabaseUserCommentRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateDatabaseUserCommentRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlDatabaseUserCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlDatabaseUserCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlDatabaseUserCommentRequest", string(data)}, " ")
}
