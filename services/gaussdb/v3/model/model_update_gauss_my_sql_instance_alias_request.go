package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGaussMySqlInstanceAliasRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID
	InstanceId string `json:"instance_id"`

	Body *ModifyAliasRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceAliasRequest", string(data)}, " ")
}
