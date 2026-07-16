package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEndpointsReq 远程接入训练作业时需要的相关配置。
type JobEndpointsReq struct {
	Ssh *SshReq `json:"ssh,omitempty"`
}

func (o JobEndpointsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEndpointsReq struct{}"
	}

	return strings.Join([]string{"JobEndpointsReq", string(data)}, " ")
}
