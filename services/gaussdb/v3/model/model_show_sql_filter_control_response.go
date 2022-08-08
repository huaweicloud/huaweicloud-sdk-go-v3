package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSqlFilterControlResponse struct {

	// sql限流开关状态。  取值：  - ON：已开启 - OFF：已关闭
	SwitchStatus   *string `json:"switch_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlFilterControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlFilterControlResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlFilterControlResponse", string(data)}, " ")
}
