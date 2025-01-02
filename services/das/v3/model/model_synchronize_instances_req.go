package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SynchronizeInstancesReq struct {

	// 数据库引擎类型。
	EngineType *string `json:"engine_type,omitempty"`
}

func (o SynchronizeInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeInstancesReq struct{}"
	}

	return strings.Join([]string{"SynchronizeInstancesReq", string(data)}, " ")
}
