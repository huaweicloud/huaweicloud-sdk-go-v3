package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像标签
type Tags struct {
	// 标签的键。最大长度127个unicode字符，key不能为空。

	Key string `json:"key"`
	// 标签的值列表。每个值最大长度255个unicode字符，如果values为空列表，则标签的值可以是任意值。值列表中的值之间为或的关系。

	Values []string `json:"values"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
