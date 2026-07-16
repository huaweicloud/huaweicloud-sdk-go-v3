package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobEnginesResponse Response Object
type ShowTrainingJobEnginesResponse struct {

	// 训练作业引擎规格总数。
	Total *int32 `json:"total,omitempty"`

	// 引擎规格参数列表。
	Items          *[]ListEnginesItems `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowTrainingJobEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobEnginesResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobEnginesResponse", string(data)}, " ")
}
