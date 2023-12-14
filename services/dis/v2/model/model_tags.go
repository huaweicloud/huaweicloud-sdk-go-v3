package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tags struct {

	// 键。  - 不能为空。  - 对于同一资源键值唯一。  - 字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Key *string `json:"key,omitempty"`

	// 标签值列表。  如果values为空列表，则表示any_value。value之间为或的关系。
	Values *[]string `json:"values,omitempty"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
