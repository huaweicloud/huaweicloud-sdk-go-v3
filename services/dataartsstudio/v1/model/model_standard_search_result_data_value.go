package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandardSearchResultDataValue 返回的数据标准详细信息。
type StandardSearchResultDataValue struct {

	// StandElementValueVO数组。
	Records *[]StandElementValueVoList `json:"records,omitempty"`

	// 数据标准的总数。
	Total *int32 `json:"total,omitempty"`
}

func (o StandardSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandardSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"StandardSearchResultDataValue", string(data)}, " ")
}
