package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementFieldVoList 属性列表。
type StandElementFieldVoList struct {

	// 属性信息。
	Fields *[]StandElementFieldVo `json:"fields,omitempty"`
}

func (o StandElementFieldVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementFieldVoList struct{}"
	}

	return strings.Join([]string{"StandElementFieldVoList", string(data)}, " ")
}
