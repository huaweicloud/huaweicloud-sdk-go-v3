package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReconfigureExtDataSourceAction 更新数据源配置
type ReconfigureExtDataSourceAction struct {

	// 重启。
	Reboot *bool `json:"reboot,omitempty"`

	// 委托。
	Agency *string `json:"agency,omitempty"`
}

func (o ReconfigureExtDataSourceAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReconfigureExtDataSourceAction struct{}"
	}

	return strings.Join([]string{"ReconfigureExtDataSourceAction", string(data)}, " ")
}
