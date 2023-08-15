package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIterationHistoriesResponse Response Object
type ListIterationHistoriesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 历史记录
	Histories      *[]IterationHistory `json:"histories,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListIterationHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIterationHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListIterationHistoriesResponse", string(data)}, " ")
}
