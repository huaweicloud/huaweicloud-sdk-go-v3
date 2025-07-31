package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubeConfigInfo kubeconfig文件
type KubeConfigInfo struct {
}

func (o KubeConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubeConfigInfo struct{}"
	}

	return strings.Join([]string{"KubeConfigInfo", string(data)}, " ")
}
