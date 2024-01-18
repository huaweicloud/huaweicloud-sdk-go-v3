package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigrateServerReq 批量迁移服务器请求。
type BatchMigrateServerReq struct {

	// 应用服务器id集合。
	ServerIds []string `json:"server_ids"`

	// 目标云办公主机id。
	HostId string `json:"host_id"`
}

func (o BatchMigrateServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateServerReq struct{}"
	}

	return strings.Join([]string{"BatchMigrateServerReq", string(data)}, " ")
}
