package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAlarmResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmResourcesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmResourcesResponse", string(data)}, " ")
}
