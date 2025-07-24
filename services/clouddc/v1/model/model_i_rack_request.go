package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IRackRequest struct {

	// 机柜描述，大小不超过512字节
	Description *string `json:"description,omitempty"`
}

func (o IRackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IRackRequest struct{}"
	}

	return strings.Join([]string{"IRackRequest", string(data)}, " ")
}
