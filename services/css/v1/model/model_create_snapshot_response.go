package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSnapshotResponse struct {
	// 快照名称。

	Name *string `json:"name,omitempty"`
	// 快照描述。

	Description *string `json:"description,omitempty"`
	// 指定要备份的索引名称。

	Indices        *string `json:"indices,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotResponse", string(data)}, " ")
}
