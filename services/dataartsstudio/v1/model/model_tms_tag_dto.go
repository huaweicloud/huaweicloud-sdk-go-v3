package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagDto tms标签
type TmsTagDto struct {

	// key值
	Key string `json:"key"`

	// value值
	Value string `json:"value"`
}

func (o TmsTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagDto struct{}"
	}

	return strings.Join([]string{"TmsTagDto", string(data)}, " ")
}
