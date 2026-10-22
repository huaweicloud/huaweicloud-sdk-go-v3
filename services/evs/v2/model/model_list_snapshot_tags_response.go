package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotTagsResponse Response Object
type ListSnapshotTagsResponse struct {

	// **参数解释** 资源的所有标签信息。 **取值范围** 不涉及。
	Tags           *[]ListResourceTagsResponseBodyTags `json:"tags,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListSnapshotTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotTagsResponse", string(data)}, " ")
}
