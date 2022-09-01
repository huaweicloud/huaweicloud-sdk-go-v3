package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 预检查失败子任务信息体
type PrecheckFailSubJobVo struct {

	// 预检查失败子任务id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 预检查失败子任务名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 检查结果。
	CheckResult *string `json:"check_result,omitempty" xml:"check_result"`
}

func (o PrecheckFailSubJobVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckFailSubJobVo struct{}"
	}

	return strings.Join([]string{"PrecheckFailSubJobVo", string(data)}, " ")
}
