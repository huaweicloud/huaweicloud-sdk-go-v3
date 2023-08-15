package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerTags
type ServerTags struct {

	// 键。  - 最大长度127个unicode字符。  - key不能为空。
	Key string `json:"key"`

	// 值列表。  - 最多10个value。  - value不允许重复。  - 每个值最大长度255个unicode字符。  - 如果values为空则表示any_value。  - value之间为或的关系。
	Values []string `json:"values"`
}

func (o ServerTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerTags struct{}"
	}

	return strings.Join([]string{"ServerTags", string(data)}, " ")
}
