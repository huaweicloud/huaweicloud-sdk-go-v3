package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddAlarmResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAlarmResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmResourcesResponse struct{}"
	}

	return strings.Join([]string{"AddAlarmResourcesResponse", string(data)}, " ")
}
