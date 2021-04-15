package model

import (
	"encoding/json"

	"strings"
)

//
type ListTag struct {
	// 功能说明：标签键 约束：key不能为空

	Key string `json:"key"`
	// 功能描述：标签值列表。 如果values为空列表，则表示any_value。value之间为或的关系。

	Values []string `json:"values"`
}

func (o ListTag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTag struct{}"
	}

	return strings.Join([]string{"ListTag", string(data)}, " ")
}
