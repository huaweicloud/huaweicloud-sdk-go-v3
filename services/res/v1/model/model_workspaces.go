package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Workspaces struct {

	// 创建时间。
	CreatedAt int64 `json:"created_at" xml:"created_at"`

	// 描述。
	Description string `json:"description" xml:"description"`

	// 企业项目id。
	EnterpriseProjectId string `json:"enterprise_project_id" xml:"enterprise_project_id"`

	// 企业项目名称。
	EnterpriseProjectName string `json:"enterprise_project_name" xml:"enterprise_project_name"`

	// 工作空间id。
	Id string `json:"id" xml:"id"`

	// 工作空间名称。
	Name string `json:"name" xml:"name"`

	// 创建者。
	Owner string `json:"owner" xml:"owner"`

	// 状态。
	Status string `json:"status" xml:"status"`

	// 更新时间。
	UpdateAt int64 `json:"update_at" xml:"update_at"`

	// 用户id。
	UserId string `json:"userId" xml:"userId"`

	// 调用账户的项目Id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`
}

func (o Workspaces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workspaces struct{}"
	}

	return strings.Join([]string{"Workspaces", string(data)}, " ")
}
