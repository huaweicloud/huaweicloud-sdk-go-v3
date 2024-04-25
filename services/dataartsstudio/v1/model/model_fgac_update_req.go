package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FgacUpdateReq struct {

	// 细粒度认证数据开发连接列表
	FgacIds *[]FgacSingleUpdateReq `json:"fgac_ids,omitempty"`
}

func (o FgacUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FgacUpdateReq struct{}"
	}

	return strings.Join([]string{"FgacUpdateReq", string(data)}, " ")
}
