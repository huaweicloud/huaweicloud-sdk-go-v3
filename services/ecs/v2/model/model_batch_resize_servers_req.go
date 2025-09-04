package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResizeServersReq struct {
	Resize *BatchResizeServersOption `json:"resize"`
}

func (o BatchResizeServersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeServersReq struct{}"
	}

	return strings.Join([]string{"BatchResizeServersReq", string(data)}, " ")
}
