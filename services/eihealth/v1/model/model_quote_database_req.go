package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuoteDatabaseReq 导入实例请求体
type QuoteDatabaseReq struct {

	// 源项目id
	SourceProjectId string `json:"source_project_id"`

	// 源数据库列表
	SourceDatabases []DatabaseSrcReq `json:"source_databases"`
}

func (o QuoteDatabaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuoteDatabaseReq struct{}"
	}

	return strings.Join([]string{"QuoteDatabaseReq", string(data)}, " ")
}
