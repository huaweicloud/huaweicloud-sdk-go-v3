package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateTrainingJobNameResponse Response Object
type ValidateTrainingJobNameResponse struct {

	// 训练作业名称是否重复，true表示已存在，false表示不存在。
	IsDuplicate    *bool `json:"is_duplicate,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateTrainingJobNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateTrainingJobNameResponse struct{}"
	}

	return strings.Join([]string{"ValidateTrainingJobNameResponse", string(data)}, " ")
}
