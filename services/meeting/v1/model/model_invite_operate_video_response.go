package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteOperateVideoResponse Response Object
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
