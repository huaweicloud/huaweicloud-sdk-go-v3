package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpBaseResponseInfo struct {
}

func (o HttpBaseResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpBaseResponseInfo struct{}"
	}

	return strings.Join([]string{"HttpBaseResponseInfo", string(data)}, " ")
}
