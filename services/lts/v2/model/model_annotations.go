package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Annotations struct {

	// 告警列表详情
	Message string `json:"message"`

	// 日志组/流id,名称
	LogInfo string `json:"log_info"`

	// 当前值
	CurrentValue string `json:"current_value"`

	// (sql/关键词)告警详情原始数据
	OldAnnotations string `json:"old_annotations"`
}

func (o Annotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Annotations struct{}"
	}

	return strings.Join([]string{"Annotations", string(data)}, " ")
}
