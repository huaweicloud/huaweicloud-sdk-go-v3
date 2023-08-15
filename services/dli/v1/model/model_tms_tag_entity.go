package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagEntity
type TmsTagEntity struct {

	// 标签的键
	Key string `json:"key"`

	// 标签的值
	Value string `json:"value"`
}

func (o TmsTagEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagEntity struct{}"
	}

	return strings.Join([]string{"TmsTagEntity", string(data)}, " ")
}
