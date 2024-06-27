package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserQosReqInfo QOS请求用户信息
type UserQosReqInfo struct {

	// 用户pid
	Pid *string `json:"pid,omitempty"`

	// 用户接入媒体类型
	AccessMediaType *string `json:"accessMediaType,omitempty"`
}

func (o UserQosReqInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserQosReqInfo struct{}"
	}

	return strings.Join([]string{"UserQosReqInfo", string(data)}, " ")
}
