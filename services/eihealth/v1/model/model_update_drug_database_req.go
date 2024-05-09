package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugDatabaseReq 数据库更新请求体
type UpdateDrugDatabaseReq struct {

	// 是否共享
	Shareable bool `json:"shareable"`
}

func (o UpdateDrugDatabaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugDatabaseReq struct{}"
	}

	return strings.Join([]string{"UpdateDrugDatabaseReq", string(data)}, " ")
}
