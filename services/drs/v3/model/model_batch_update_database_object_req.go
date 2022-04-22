package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量更新数据库对象选择消息体
type BatchUpdateDatabaseObjectReq struct {

	// 批量更新数据库对象请求任务ID列表
	Jobs []UpdateDatabaseObjectReq `json:"jobs"`
}

func (o BatchUpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateDatabaseObjectReq", string(data)}, " ")
}
