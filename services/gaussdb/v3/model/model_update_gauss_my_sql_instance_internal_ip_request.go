package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceInternalIpRequest Request Object
type UpdateGaussMySqlInstanceInternalIpRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyInternalIpRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceInternalIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceInternalIpRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceInternalIpRequest", string(data)}, " ")
}
