package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseDto 包含用户工作相关信息的对象
type EnterpriseDto struct {

	// 成本中心
	CostCenter *string `json:"costCenter,omitempty"`

	// 部门
	Department *string `json:"department,omitempty"`

	// 分部
	Division *string `json:"division,omitempty"`

	// 员工编号
	EmployeeNumber *string `json:"employeeNumber,omitempty"`

	Manager *ManagerDto `json:"manager,omitempty"`

	// 组织
	Organization *string `json:"organization,omitempty"`
}

func (o EnterpriseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseDto struct{}"
	}

	return strings.Join([]string{"EnterpriseDto", string(data)}, " ")
}
