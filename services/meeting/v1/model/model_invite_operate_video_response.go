package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type InviteOperateVideoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InviteOperateVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteOperateVideoResponse struct{}"
	}

	return strings.Join([]string{"InviteOperateVideoResponse", string(data)}, " ")
}
