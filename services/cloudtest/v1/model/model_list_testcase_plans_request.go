package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestcasePlansRequest Request Object
type ListTestcasePlansRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 版本URI
	BranchUri string `json:"branch_uri"`

	Body *TestcasePlanQueryParam `json:"body,omitempty"`
}

func (o ListTestcasePlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestcasePlansRequest struct{}"
	}

	return strings.Join([]string{"ListTestcasePlansRequest", string(data)}, " ")
}
