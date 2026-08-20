package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateSnapshotDeletableVo 批量更新快照可删除标识请求对象。
type BatchUpdateSnapshotDeletableVo struct {

	// 快照ID列表。通过接口查询工作项计划管理快照列表获取响应参数中的id字段。
	Ids []string `json:"ids"`

	// 是否为可删除标识。
	Deletable bool `json:"deletable"`
}

func (o BatchUpdateSnapshotDeletableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSnapshotDeletableVo struct{}"
	}

	return strings.Join([]string{"BatchUpdateSnapshotDeletableVo", string(data)}, " ")
}
