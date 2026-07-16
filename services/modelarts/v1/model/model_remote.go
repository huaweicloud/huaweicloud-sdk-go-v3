package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Remote 数据实际输出信息。
type Remote struct {
	Obs *RemoteObs `json:"obs"`
}

func (o Remote) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Remote struct{}"
	}

	return strings.Join([]string{"Remote", string(data)}, " ")
}
