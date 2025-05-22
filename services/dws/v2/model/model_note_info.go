package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NoteInfo **参数解释**： 重分布提示信息。 **取值范围**： 不涉及。
type NoteInfo struct {
	BucketSplitInfo *BucketSplitInfo `json:"bucket_split_info,omitempty"`
}

func (o NoteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoteInfo struct{}"
	}

	return strings.Join([]string{"NoteInfo", string(data)}, " ")
}
