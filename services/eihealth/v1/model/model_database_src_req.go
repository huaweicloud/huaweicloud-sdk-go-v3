package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseSrcReq struct {

	// 源数据库id
	SourceDatabaseId string `json:"source_database_id"`
}

func (o DatabaseSrcReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseSrcReq struct{}"
	}

	return strings.Join([]string{"DatabaseSrcReq", string(data)}, " ")
}
