package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnhancedConnectionPrivilege 增强型跨源连接各个授权项目的信息。
type EnhancedConnectionPrivilege struct {

	// 授权时object的信息。
	Object *string `json:"object,omitempty"`

	// 授权的项目ID。
	ApplicantProjectId *string `json:"applicant_project_id,omitempty"`

	// 授权操作信息。
	Privileges *[]string `json:"privileges,omitempty"`
}

func (o EnhancedConnectionPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionPrivilege struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionPrivilege", string(data)}, " ")
}
