package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberGroupandStreamLogStreams struct {
	LogStreamName *string `json:"log_stream_name,omitempty"`

	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o MemberGroupandStreamLogStreams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberGroupandStreamLogStreams struct{}"
	}

	return strings.Join([]string{"MemberGroupandStreamLogStreams", string(data)}, " ")
}
