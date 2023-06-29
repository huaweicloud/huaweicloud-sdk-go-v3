package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDisasterRecoveryDrillRequestParams 创建容灾演练请求数据结构
type CreateDisasterRecoveryDrillRequestParams struct {

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`

	// 演练虚拟私有云ID，不指定时系统会自动创建演练VPC。
	DrillVpcId *string `json:"drill_vpc_id,omitempty"`

	// 指定容灾演练的名称，最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。
	Name string `json:"name"`
}

func (o CreateDisasterRecoveryDrillRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillRequestParams struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillRequestParams", string(data)}, " ")
}
