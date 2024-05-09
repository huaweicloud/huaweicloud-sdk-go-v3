package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDrugDatabaseFileReq 追加数据库文件请求体
type AddDrugDatabaseFileReq struct {
	File *DatabaseFile `json:"file"`

	// 数据库描述
	Description *string `json:"description,omitempty"`
}

func (o AddDrugDatabaseFileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDrugDatabaseFileReq struct{}"
	}

	return strings.Join([]string{"AddDrugDatabaseFileReq", string(data)}, " ")
}
