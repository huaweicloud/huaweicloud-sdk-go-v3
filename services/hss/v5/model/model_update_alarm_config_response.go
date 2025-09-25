package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmConfigResponse Response Object
type UpdateAlarmConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmConfigResponse", string(data)}, " ")
}
