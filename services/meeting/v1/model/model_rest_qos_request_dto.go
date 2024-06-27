package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestQosRequestDto QOS请求
type RestQosRequestDto struct {

	// 用户pid 媒体接入类型列表
	Users *[]UserQosReqInfo `json:"users,omitempty"`
}

func (o RestQosRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestQosRequestDto struct{}"
	}

	return strings.Join([]string{"RestQosRequestDto", string(data)}, " ")
}
