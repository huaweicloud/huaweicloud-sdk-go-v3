package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExercisesResponse Response Object
type ListExercisesResponse struct {

	// 习题库数量
	Total *int32 `json:"total,omitempty"`

	// 习题库列表
	Data           *[]PackageExerciseCard `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListExercisesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExercisesResponse struct{}"
	}

	return strings.Join([]string{"ListExercisesResponse", string(data)}, " ")
}
