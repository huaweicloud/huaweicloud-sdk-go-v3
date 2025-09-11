package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportAccountSessionNew struct {

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据库用户
	DbUser *string `json:"db_user,omitempty"`

	// SESSION数量
	SessionNum *int64 `json:"session_num,omitempty"`
}

func (o ReportAccountSessionNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportAccountSessionNew struct{}"
	}

	return strings.Join([]string{"ReportAccountSessionNew", string(data)}, " ")
}
