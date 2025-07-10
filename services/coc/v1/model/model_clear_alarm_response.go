package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearAlarmResponse Response Object
type ClearAlarmResponse struct {

	// 返回数据
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ClearAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAlarmResponse struct{}"
	}

	return strings.Join([]string{"ClearAlarmResponse", string(data)}, " ")
}
