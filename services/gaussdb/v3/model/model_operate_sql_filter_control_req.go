package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateSqlFilterControlReq 开启/关闭SQL限流参数体
type OperateSqlFilterControlReq struct {

	// SQL限流开关状态。 取值： - ON，表示开启。 - OFF，表示关闭。
	SwitchStatus string `json:"switch_status"`
}

func (o OperateSqlFilterControlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSqlFilterControlReq struct{}"
	}

	return strings.Join([]string{"OperateSqlFilterControlReq", string(data)}, " ")
}
