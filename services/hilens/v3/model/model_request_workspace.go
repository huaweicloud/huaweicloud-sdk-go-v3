package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestWorkspace struct {

	// 工作空间名
	Name string `json:"name"`

	// 企业项目id，默认为0
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 企业项目名，默认为default
	EnterpriseProjectName string `json:"enterprise_project_name"`

	// 工作空间描述
	Description *string `json:"description,omitempty"`
}

func (o RequestWorkspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestWorkspace struct{}"
	}

	return strings.Join([]string{"RequestWorkspace", string(data)}, " ")
}
