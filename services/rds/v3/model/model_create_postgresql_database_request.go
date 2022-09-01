package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePostgresqlDatabaseRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *PostgresqlDatabaseForCreation `json:"body,omitempty" xml:"body"`
}

func (o CreatePostgresqlDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseRequest", string(data)}, " ")
}
