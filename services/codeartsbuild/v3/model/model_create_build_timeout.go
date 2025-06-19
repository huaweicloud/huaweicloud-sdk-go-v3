package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildTimeout 构建任务超时时间
type CreateBuildTimeout struct {

	// 超时时间数值
	Limit *string `json:"limit,omitempty"`

	// 超时时间单位
	Unit *string `json:"unit,omitempty"`
}

func (o CreateBuildTimeout) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildTimeout struct{}"
	}

	return strings.Join([]string{"CreateBuildTimeout", string(data)}, " ")
}
