package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelRelationsResultDataValueRecords 数据记录。
type ListTableModelRelationsResultDataValueRecords struct {

	// TableModelVO信息。
	Tables *[]TableModelVo `json:"tables,omitempty"`

	// 层级信息信息。
	Inheritances *[]interface{} `json:"inheritances,omitempty"`

	// RelationVO信息。
	Relations *[]RelationVo `json:"relations,omitempty"`
}

func (o ListTableModelRelationsResultDataValueRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsResultDataValueRecords struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsResultDataValueRecords", string(data)}, " ")
}
