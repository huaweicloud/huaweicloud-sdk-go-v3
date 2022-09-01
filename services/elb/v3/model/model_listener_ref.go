package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监听器
type ListenerRef struct {

	// 监听器ID。
	Id string `json:"id" xml:"id"`
}

func (o ListenerRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerRef struct{}"
	}

	return strings.Join([]string{"ListenerRef", string(data)}, " ")
}
