package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmRemote 数据实际输入信息，异构作业只支持OBS。
type AlgorithmRemote struct {
	Obs *RemoteObs `json:"obs,omitempty"`
}

func (o AlgorithmRemote) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmRemote struct{}"
	}

	return strings.Join([]string{"AlgorithmRemote", string(data)}, " ")
}
