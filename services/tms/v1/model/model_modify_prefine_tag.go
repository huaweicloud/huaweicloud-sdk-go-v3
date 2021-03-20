package model

import (
	"encoding/json"

	"strings"
)

// 修改预定义标签
type ModifyPrefineTag struct {
	NewTag *PredefineTagRequest `json:"new_tag"`

	OldTag *PredefineTagRequest `json:"old_tag"`
}

func (o ModifyPrefineTag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyPrefineTag struct{}"
	}

	return strings.Join([]string{"ModifyPrefineTag", string(data)}, " ")
}
