package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartInstanceRsp 重启集群返回值
type RestartInstanceRsp struct {
	JobId *[]string `json:"jobId,omitempty"`
}

func (o RestartInstanceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRsp struct{}"
	}

	return strings.Join([]string{"RestartInstanceRsp", string(data)}, " ")
}
