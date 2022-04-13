package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchDepartmentByNameResponse struct {
	// 部门信息

	Body           *[]QueryDeptResultDto `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o SearchDepartmentByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDepartmentByNameResponse struct{}"
	}

	return strings.Join([]string{"SearchDepartmentByNameResponse", string(data)}, " ")
}
