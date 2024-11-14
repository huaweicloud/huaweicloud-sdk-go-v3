package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceScheduleWindowResponse Response Object
type CancelInstanceScheduleWindowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelInstanceScheduleWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceScheduleWindowResponse struct{}"
	}

	return strings.Join([]string{"CancelInstanceScheduleWindowResponse", string(data)}, " ")
}
