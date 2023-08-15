package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceSecurityGroupRequest Request Object
type UpdateGaussMySqlInstanceSecurityGroupRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifySecurityGroupRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceSecurityGroupRequest", string(data)}, " ")
}
