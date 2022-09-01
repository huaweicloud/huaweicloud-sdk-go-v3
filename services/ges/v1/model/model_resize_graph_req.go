package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeGraphReq struct {
	Resize *GraphSizeTypeIndexReq `json:"resize" xml:"resize"`
}

func (o ResizeGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeGraphReq struct{}"
	}

	return strings.Join([]string{"ResizeGraphReq", string(data)}, " ")
}
