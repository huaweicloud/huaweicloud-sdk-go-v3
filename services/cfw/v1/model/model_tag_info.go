package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagInfo struct {

	// **参数解释**： 标签键 **约束限制**： 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值列表 **约束限制**： 不涉及
	Values *[]string `json:"values,omitempty"`
}

func (o TagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfo struct{}"
	}

	return strings.Join([]string{"TagInfo", string(data)}, " ")
}
