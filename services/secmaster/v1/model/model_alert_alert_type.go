package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertAlertType 告警分类，详细定义参考《告警类型定义》
type AlertAlertType struct {

	// 类别
	Category *string `json:"category,omitempty"`

	// 告警类型
	AlertType *string `json:"alert_type,omitempty"`
}

func (o AlertAlertType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertAlertType struct{}"
	}

	return strings.Join([]string{"AlertAlertType", string(data)}, " ")
}
