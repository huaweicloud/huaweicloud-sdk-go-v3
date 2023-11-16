package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpStatusDetail 封堵信息
type IpStatusDetail struct {

	// 封堵时间
	BlockTime int64 `json:"block_time"`

	// 解封时间
	UnblockTime int64 `json:"unblock_time"`
}

func (o IpStatusDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpStatusDetail struct{}"
	}

	return strings.Join([]string{"IpStatusDetail", string(data)}, " ")
}
