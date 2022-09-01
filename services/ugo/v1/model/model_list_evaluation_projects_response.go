package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEvaluationProjectsResponse struct {

	// 评估项目总数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 当前页的评估项目列表。
	EvaluationProjects *[]EvaluationProject `json:"evaluation_projects,omitempty" xml:"evaluation_projects"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListEvaluationProjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvaluationProjectsResponse struct{}"
	}

	return strings.Join([]string{"ListEvaluationProjectsResponse", string(data)}, " ")
}
