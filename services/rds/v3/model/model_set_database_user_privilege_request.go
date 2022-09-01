package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetDatabaseUserPrivilegeRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *SetDatabaseUserPrivilegeReqV3 `json:"body,omitempty" xml:"body"`
}

func (o SetDatabaseUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDatabaseUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"SetDatabaseUserPrivilegeRequest", string(data)}, " ")
}
