package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteResp 数据实际输出信息。
type RemoteResp struct {
	Obs *RemoteObsResp `json:"obs"`
}

func (o RemoteResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteResp struct{}"
	}

	return strings.Join([]string{"RemoteResp", string(data)}, " ")
}
