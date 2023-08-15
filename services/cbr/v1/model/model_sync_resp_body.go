package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncRespBody struct {

	// 是否自动触发
	OperationLogId string `json:"operation_log_id"`

	// 混合云vault id
	VaultId string `json:"vault_id"`
}

func (o SyncRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncRespBody struct{}"
	}

	return strings.Join([]string{"SyncRespBody", string(data)}, " ")
}
