package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtocolResVo struct {

	// 期望响应的状态码
	ResponseCode *[]int32 `json:"response_code,omitempty"`

	// 期望响应时间
	ResponseTime *string `json:"response_time,omitempty"`
}

func (o ProtocolResVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtocolResVo struct{}"
	}

	return strings.Join([]string{"ProtocolResVo", string(data)}, " ")
}
