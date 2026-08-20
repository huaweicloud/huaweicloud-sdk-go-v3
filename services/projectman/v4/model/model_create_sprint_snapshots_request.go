package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSprintSnapshotsRequest Request Object
type CreateSprintSnapshotsRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	Body *SprintSnapshotsCreateParam `json:"body,omitempty"`
}

func (o CreateSprintSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSprintSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"CreateSprintSnapshotsRequest", string(data)}, " ")
}
