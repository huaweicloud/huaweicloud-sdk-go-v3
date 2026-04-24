package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchFactoryBaselinesResponse Response Object
type ListSearchFactoryBaselinesResponse struct {

	// 基线任务。
	Baselines *[]BaselineV2 `json:"baselines,omitempty"`

	// 总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSearchFactoryBaselinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchFactoryBaselinesResponse struct{}"
	}

	return strings.Join([]string{"ListSearchFactoryBaselinesResponse", string(data)}, " ")
}
