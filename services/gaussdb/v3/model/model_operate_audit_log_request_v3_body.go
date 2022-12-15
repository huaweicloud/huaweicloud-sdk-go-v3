package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 开启/关闭全量SQL参数体
type OperateAuditLogRequestV3Body struct {

	// 全量SQL开关状态。 取值： - ON，表示开启 - OFF，表示关闭
	SwitchStatus string `json:"switch_status"`
}

func (o OperateAuditLogRequestV3Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateAuditLogRequestV3Body struct{}"
	}

	return strings.Join([]string{"OperateAuditLogRequestV3Body", string(data)}, " ")
}
