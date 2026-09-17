package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineSnapshotsRequest Request Object
type ShowBaselineSnapshotsRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 特性集快照版本ID，不传则查询当前版本特性集，传值则查询对应版本的特性集
	SnapshotVersionId *string `json:"snapshot_version_id,omitempty"`
}

func (o ShowBaselineSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaselineSnapshotsRequest", string(data)}, " ")
}
