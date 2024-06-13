package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PiInfo pi迭代筛选条件
type PiInfo struct {

	// 迭代列表
	Sprints *[]string `json:"sprints,omitempty"`

	// pi的id，层级关系：pi -> 迭代 -> 需求
	PiId *string `json:"pi_id,omitempty"`
}

func (o PiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PiInfo struct{}"
	}

	return strings.Join([]string{"PiInfo", string(data)}, " ")
}
