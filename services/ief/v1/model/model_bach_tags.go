package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BachTags struct {
	// 标签列表

	Tags *[]ResourceTag `json:"tags,omitempty"`
	// 操作标识：仅限于create（创建）、delete（删除）

	Action string `json:"action"`
}

func (o BachTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BachTags struct{}"
	}

	return strings.Join([]string{"BachTags", string(data)}, " ")
}
