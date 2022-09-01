package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeSqlSwitchResponse struct {

	// 开关状态。取值： Enabled：已开启， Disabled：已关闭， Switching：开关切换中
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeSqlSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlSwitchResponse struct{}"
	}

	return strings.Join([]string{"ChangeSqlSwitchResponse", string(data)}, " ")
}
