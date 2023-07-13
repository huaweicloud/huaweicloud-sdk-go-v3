package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIpsSwitchUsingPostResponse Response Object
type ChangeIpsSwitchUsingPostResponse struct {

	// trace_id
	TraceId *string `json:"trace_id,omitempty"`

	Data           *CommonResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ChangeIpsSwitchUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsSwitchUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ChangeIpsSwitchUsingPostResponse", string(data)}, " ")
}
