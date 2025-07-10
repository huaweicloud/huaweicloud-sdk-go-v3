package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesOfPolicyGroupRequest Request Object
type ListPoliciesOfPolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	// 根据策略类型过滤结果，不传则查询所有策略。 可选类型: - 外设：Peripherals; - 音频：Audio; - 客户端：Client; - 显示：Display; - 文件与剪切板：FileAndClip; - 接入控制：ClientAccessControl; - 会话：SessionAutoDisconnect; - 虚拟通道：VirtualChannel - 水印：Watermark; - 键盘鼠标：KeyboardAndMouse; - 通用音视频旁路：Seamless。 - 录屏审计：RecordAudit。
	PolicyType *string `json:"policy_type,omitempty"`
}

func (o ListPoliciesOfPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesOfPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"ListPoliciesOfPolicyGroupRequest", string(data)}, " ")
}
