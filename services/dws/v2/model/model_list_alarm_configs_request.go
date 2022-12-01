package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmConfigsRequest struct {
}

func (o ListAlarmConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmConfigsRequest", string(data)}, " ")
}
