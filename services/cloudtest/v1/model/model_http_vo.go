package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpVo struct {
	Request *ProtocolReqVo `json:"request,omitempty"`

	Response *ProtocolResVo `json:"response,omitempty"`
}

func (o HttpVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpVo struct{}"
	}

	return strings.Join([]string{"HttpVo", string(data)}, " ")
}
