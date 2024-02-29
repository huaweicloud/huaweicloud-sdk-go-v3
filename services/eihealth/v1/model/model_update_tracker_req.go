package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrackerReq 更新项目追踪器
type UpdateTrackerReq struct {

	// 审计数据类型列表
	DataEvent []DataEvent `json:"data_event"`
}

func (o UpdateTrackerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerReq struct{}"
	}

	return strings.Join([]string{"UpdateTrackerReq", string(data)}, " ")
}
