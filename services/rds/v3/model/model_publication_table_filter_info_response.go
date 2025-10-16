package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicationTableFilterInfoResponse 表筛选器。
type PublicationTableFilterInfoResponse struct {

	// 筛选关系。  为空时表示当前为最下级筛选器，不为空时表示当前项还有下级筛选器。
	Relation *string `json:"relation,omitempty"`

	// 筛选字段。
	Column *string `json:"column,omitempty"`

	// 筛选条件。
	Condition *string `json:"condition,omitempty"`

	// 筛选值。
	Value *string `json:"value,omitempty"`

	// 下级筛选器。
	Filters *[]PublicationTableFilterInfoResponse `json:"filters,omitempty"`
}

func (o PublicationTableFilterInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicationTableFilterInfoResponse struct{}"
	}

	return strings.Join([]string{"PublicationTableFilterInfoResponse", string(data)}, " ")
}
