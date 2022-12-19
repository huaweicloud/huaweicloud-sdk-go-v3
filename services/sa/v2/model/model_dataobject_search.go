package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Dataobject Search
type DataobjectSearch struct {

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// sortby
	SortBy *string `json:"sort_by,omitempty"`

	// order
	Order *string `json:"order,omitempty"`

	// search start time
	FromDate *string `json:"from_date,omitempty"`

	// search end time
	ToDate *string `json:"to_date,omitempty"`

	Condition *DataobjectSearchCondition `json:"condition,omitempty"`
}

func (o DataobjectSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearch struct{}"
	}

	return strings.Join([]string{"DataobjectSearch", string(data)}, " ")
}
