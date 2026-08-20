package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotResult 创建快照的结果
type CreateSnapshotResult struct {

	// 快照ID。
	Id *string `json:"id,omitempty"`

	// 快照名称。创建时自动生成，工作项快照名称生成规则为：“工作项类型”+“ v” + “年”+“.”+“月”+“.”+“日”+“.”+“当日生成版本次数”。例如工作项类型为IR的工作项在2026年3月25日第一次打快照系统生成的快照名称为：IR v26.03.25.1。
	Title *string `json:"title,omitempty"`

	// 快照类型。工作项快照固定为：issue_snap_item。
	Category *string `json:"category,omitempty"`

	// 快照的工作项ID。
	IssueId *string `json:"issue_id,omitempty"`

	// 快照的创建人ID。
	CreatedBy *string `json:"created_by,omitempty"`

	// 快照是否可被删除。
	Deletable *bool `json:"deletable,omitempty"`

	// 创建快照失败的原因。
	Errormsg *string `json:"errormsg,omitempty"`
}

func (o CreateSnapshotResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotResult struct{}"
	}

	return strings.Join([]string{"CreateSnapshotResult", string(data)}, " ")
}
