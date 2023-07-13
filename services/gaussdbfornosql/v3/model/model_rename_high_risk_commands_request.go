package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameHighRiskCommandsRequest 修改高危命令的请求体
type RenameHighRiskCommandsRequest struct {

	// 高危命令与对应重命名命令。
	Commands []CommandInfo `json:"commands"`
}

func (o RenameHighRiskCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameHighRiskCommandsRequest struct{}"
	}

	return strings.Join([]string{"RenameHighRiskCommandsRequest", string(data)}, " ")
}
