package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsDbRequest struct {

	// 数据库列表
	Databases []RdsDbRequestDatabases `json:"databases"`
}

func (o RdsDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsDbRequest struct{}"
	}

	return strings.Join([]string{"RdsDbRequest", string(data)}, " ")
}
