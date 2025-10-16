package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicationTableFilterInfoRequest 表筛选器。
type PublicationTableFilterInfoRequest struct {

	// 筛选关系。  为空时表示当前为最下级筛选器，不为空时表示当前项还有下级筛选器 不为空时取值范围：AND, OR
	Relation *string `json:"relation,omitempty"`

	// 筛选字段（仅当筛选关系为空时生效，且筛选关系为空时必传）。
	Column *string `json:"column,omitempty"`

	// 筛选条件（仅当筛选关系为空时生效，且筛选关系为空时必传）。 取值范围： =, !=, >, <, >=, <=, LIKE, NOT LIKE, IN
	Condition *string `json:"condition,omitempty"`

	// 筛选值（仅当筛选关系为空时生效，且筛选关系为空时必传）。
	Value *string `json:"value,omitempty"`

	// 下级筛选器（仅当筛选关系不为空时生效，且筛选关系不为空时必传）。
	Filters *[]PublicationTableFilterInfoRequest `json:"filters,omitempty"`
}

func (o PublicationTableFilterInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationTableFilterInfoRequest struct{}"
	}

	return strings.Join([]string{"PublicationTableFilterInfoRequest", string(data)}, " ")
}
