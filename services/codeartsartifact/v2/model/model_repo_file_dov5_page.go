package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoFileDov5Page struct {

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalRecords *int32 `json:"total_records,omitempty"`

	// **参数解释**: 总页数。 **取值范围**: 不涉及。
	TotalPages *int32 `json:"total_pages,omitempty"`

	// **参数解释**: 文件列表。 **取值范围**: 不涉及。
	Data *[]RepoFileDov5 `json:"data,omitempty"`
}

func (o RepoFileDov5Page) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoFileDov5Page struct{}"
	}

	return strings.Join([]string{"RepoFileDov5Page", string(data)}, " ")
}
