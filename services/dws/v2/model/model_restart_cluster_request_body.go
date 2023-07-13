package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartClusterRequestBody This is a auto create Body Object
type RestartClusterRequestBody struct {

	// 重启标识
	Restart *interface{} `json:"restart"`
}

func (o RestartClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterRequestBody struct{}"
	}

	return strings.Join([]string{"RestartClusterRequestBody", string(data)}, " ")
}
