package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchAlgorithmsResponse Response Object
type ShowSearchAlgorithmsResponse struct {

	// 超参搜索算法的个数。
	SearchAlgoCount *int32 `json:"search_algo_count,omitempty"`

	// 所有超参搜索算法的列表。
	SearchAlgoList *[]ListSearchAlgorithmsSearchAlgoList `json:"search_algo_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowSearchAlgorithmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchAlgorithmsResponse struct{}"
	}

	return strings.Join([]string{"ShowSearchAlgorithmsResponse", string(data)}, " ")
}
