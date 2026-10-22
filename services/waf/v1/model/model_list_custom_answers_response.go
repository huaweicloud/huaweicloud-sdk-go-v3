package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomAnswersResponse Response Object
type ListCustomAnswersResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 代答库列表
	Items          *[]CustomAnswersItem `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListCustomAnswersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomAnswersResponse struct{}"
	}

	return strings.Join([]string{"ListCustomAnswersResponse", string(data)}, " ")
}
