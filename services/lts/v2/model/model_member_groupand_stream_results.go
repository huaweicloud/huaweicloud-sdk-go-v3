package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberGroupandStreamResults struct {
	LogGroupId *string `json:"log_group_id,omitempty"`

	LogGroupName *string `json:"log_group_name,omitempty"`

	LogStreams *[]MemberGroupandStreamLogStreams `json:"log_streams,omitempty"`
}

func (o MemberGroupandStreamResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberGroupandStreamResults struct{}"
	}

	return strings.Join([]string{"MemberGroupandStreamResults", string(data)}, " ")
}
