package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAllDifficultsResponse struct {

	// 所有难度数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 难度信息
	Data           *[]DifficultInfo `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAllDifficultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllDifficultsResponse struct{}"
	}

	return strings.Join([]string{"ListAllDifficultsResponse", string(data)}, " ")
}
