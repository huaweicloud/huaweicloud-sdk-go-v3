package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectedProblem 每项检查项检测到的问题
type DetectedProblem struct {

	// 存在问题的数据库对象
	DbObject *string `json:"db_object,omitempty"`

	// 告警级别
	Level *string `json:"level,omitempty"`

	// 检查项描述
	Description *string `json:"description,omitempty"`
}

func (o DetectedProblem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectedProblem struct{}"
	}

	return strings.Join([]string{"DetectedProblem", string(data)}, " ")
}
