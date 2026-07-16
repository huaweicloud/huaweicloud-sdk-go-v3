package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEipAlarmWhitelistResponse Response Object
type AddEipAlarmWhitelistResponse struct {
	Data           *AddEipAlarmWhitelistRespData `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o AddEipAlarmWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipAlarmWhitelistResponse struct{}"
	}

	return strings.Join([]string{"AddEipAlarmWhitelistResponse", string(data)}, " ")
}
