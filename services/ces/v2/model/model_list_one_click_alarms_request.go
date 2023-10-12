package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOneClickAlarmsRequest Request Object
type ListOneClickAlarmsRequest struct {
}

func (o ListOneClickAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmsRequest", string(data)}, " ")
}
