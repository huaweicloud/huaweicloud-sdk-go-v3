package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePostgresqlInstanceAliasRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateRdsInstanceAliasRequest `json:"body,omitempty" xml:"body"`
}

func (o UpdatePostgresqlInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlInstanceAliasRequest", string(data)}, " ")
}
