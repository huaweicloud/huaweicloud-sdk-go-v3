package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryAllSearchCriteriasResponse Response Object
type ListQueryAllSearchCriteriasResponse struct {

	// 快速查询
	SearchCriterias *[]SearchCriteriasBody `json:"search_criterias,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o ListQueryAllSearchCriteriasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryAllSearchCriteriasResponse struct{}"
	}

	return strings.Join([]string{"ListQueryAllSearchCriteriasResponse", string(data)}, " ")
}
