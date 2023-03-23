package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskModelMarketPlace struct {

	// 任务类型
	Type int32 `json:"type"`
}

func (o TaskModelMarketPlace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskModelMarketPlace struct{}"
	}

	return strings.Join([]string{"TaskModelMarketPlace", string(data)}, " ")
}
