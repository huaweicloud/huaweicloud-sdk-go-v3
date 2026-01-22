package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEipAlarmWhitelistRequest Request Object
type AddEipAlarmWhitelistRequest struct {
	Body *AddEipAlarmWhitelistRequestBody `json:"body,omitempty"`
}

func (o AddEipAlarmWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipAlarmWhitelistRequest struct{}"
	}

	return strings.Join([]string{"AddEipAlarmWhitelistRequest", string(data)}, " ")
}
