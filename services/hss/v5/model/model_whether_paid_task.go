package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WhetherPaidTask 此次扫描任务是否付费
type WhetherPaidTask struct {
}

func (o WhetherPaidTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhetherPaidTask struct{}"
	}

	return strings.Join([]string{"WhetherPaidTask", string(data)}, " ")
}
