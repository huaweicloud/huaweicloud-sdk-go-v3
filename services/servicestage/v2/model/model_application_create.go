package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationCreate struct {

	// 应用名称。
	Name string `json:"name" xml:"name"`

	// 应用描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 企业项目ID。默认值为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ApplicationCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationCreate struct{}"
	}

	return strings.Join([]string{"ApplicationCreate", string(data)}, " ")
}
