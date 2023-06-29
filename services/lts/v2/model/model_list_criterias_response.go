package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCriteriasResponse Response Object
type ListCriteriasResponse struct {
	SearchCriterias *[]GetQuerySearchCriteriasBody `json:"search_criterias,omitempty"`
	HttpStatusCode  int                            `json:"-"`
}

func (o ListCriteriasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCriteriasResponse struct{}"
	}

	return strings.Join([]string{"ListCriteriasResponse", string(data)}, " ")
}
