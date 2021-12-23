package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteActiveAlarmsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteActiveAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteActiveAlarmsResponse struct{}"
	}

	return strings.Join([]string{"DeleteActiveAlarmsResponse", string(data)}, " ")
}
