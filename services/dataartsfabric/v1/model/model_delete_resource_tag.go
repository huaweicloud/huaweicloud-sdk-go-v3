package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceTag 标签。
type DeleteResourceTag struct {

	// 键。
	Key string `json:"key"`

	// 值。
	Value *string `json:"value,omitempty"`
}

func (o DeleteResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTag struct{}"
	}

	return strings.Join([]string{"DeleteResourceTag", string(data)}, " ")
}
