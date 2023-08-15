package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyOpsWindow struct {

	// 维护起始时间，UTC时间。
	StartTime string `json:"start_time"`

	// 维护结束时间，UTC时间。  说明：GuassDB(for MySQL)数据库的结束时间和开始时间必须为整点时间，且相隔四个小时。
	EndTime string `json:"end_time"`
}

func (o ModifyOpsWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyOpsWindow struct{}"
	}

	return strings.Join([]string{"ModifyOpsWindow", string(data)}, " ")
}
