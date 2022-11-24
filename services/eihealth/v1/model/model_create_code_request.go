package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCodeRequest struct {

	// 用户id
	UserId string `json:"user_id"`

	Body *SendCodeReq `json:"body,omitempty"`
}

func (o CreateCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeRequest struct{}"
	}

	return strings.Join([]string{"CreateCodeRequest", string(data)}, " ")
}
