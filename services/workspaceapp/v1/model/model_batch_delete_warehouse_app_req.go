package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWarehouseAppReq 批量删除应用仓库中的应用请求体。
type BatchDeleteWarehouseAppReq struct {

	// 应用ID,最多同时操作10个。
	Ids []string `json:"ids"`
}

func (o BatchDeleteWarehouseAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWarehouseAppReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteWarehouseAppReq", string(data)}, " ")
}
