package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddRelationsInfo struct {
	Relations *[]RelationInfo `json:"relations,omitempty"`

	// 工作项类型id
	TrackerId *string `json:"tracker_id,omitempty"`

	// 版本uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 测试套id
	TaskUri *string `json:"task_uri,omitempty"`

	// 是否将需求添加到迭代
	AddToIterator *bool `json:"add_to_iterator,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o AddRelationsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRelationsInfo struct{}"
	}

	return strings.Join([]string{"AddRelationsInfo", string(data)}, " ")
}
