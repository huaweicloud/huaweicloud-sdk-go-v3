package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipRequestBodyGlobalEip 更新全域弹性公网IP请求体
type UpdateGlobalEipRequestBodyGlobalEip struct {

	// - 功能说明：全域弹性公网IP名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`
}

func (o UpdateGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
