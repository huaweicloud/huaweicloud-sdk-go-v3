package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataProcessReq 指定任务数据加工规则请求体
type DataProcessReq struct {

	// 指定任务数据加工规则请求体
	DataProcessInfo *[]DataProcessInfo `json:"data_process_info,omitempty"`
}

func (o DataProcessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataProcessReq struct{}"
	}

	return strings.Join([]string{"DataProcessReq", string(data)}, " ")
}
