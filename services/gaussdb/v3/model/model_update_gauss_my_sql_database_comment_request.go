package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlDatabaseCommentRequest Request Object
type UpdateGaussMySqlDatabaseCommentRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateDatabaseCommentRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlDatabaseCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlDatabaseCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlDatabaseCommentRequest", string(data)}, " ")
}
