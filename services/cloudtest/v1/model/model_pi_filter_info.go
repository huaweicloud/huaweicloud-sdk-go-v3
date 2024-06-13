package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PiFilterInfo pi筛选条件
type PiFilterInfo struct {

	// pi迭代筛选条件
	PiSprints *[]PiInfo `json:"pi_sprints,omitempty"`

	// pi下拉框全选标识，全选时为true
	AllPi *bool `json:"all_pi,omitempty"`
}

func (o PiFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PiFilterInfo struct{}"
	}

	return strings.Join([]string{"PiFilterInfo", string(data)}, " ")
}
