package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSnapshotRecordVo 批量删除快照请求对象。
type BatchDeleteSnapshotRecordVo struct {

	// 快照ID列表。通过接口查询工作项计划管理快照列表获取响应参数中的id字段。
	Ids []string `json:"ids"`
}

func (o BatchDeleteSnapshotRecordVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSnapshotRecordVo struct{}"
	}

	return strings.Join([]string{"BatchDeleteSnapshotRecordVo", string(data)}, " ")
}
