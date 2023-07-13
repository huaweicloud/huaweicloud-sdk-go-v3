package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExercisesRequest Request Object
type ListExercisesRequest struct {

	// 习题库id
	PackageId string `json:"package_id"`

	Body *ExercisesListRequestBody `json:"body,omitempty"`
}

func (o ListExercisesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExercisesRequest struct{}"
	}

	return strings.Join([]string{"ListExercisesRequest", string(data)}, " ")
}
