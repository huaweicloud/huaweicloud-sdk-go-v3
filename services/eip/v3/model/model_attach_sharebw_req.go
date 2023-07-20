package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachSharebwReq 共享带宽加入弹性公网IP请求参数
type AttachSharebwReq struct {
	Publicip *AttachSharebwDict `json:"publicip,omitempty"`
}

func (o AttachSharebwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSharebwReq struct{}"
	}

	return strings.Join([]string{"AttachSharebwReq", string(data)}, " ")
}
