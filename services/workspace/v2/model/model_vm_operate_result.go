package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VmOperateResult 操作失败的桌面列表。
type VmOperateResult struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 操作失败的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 操作失败的原因描述。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o VmOperateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VmOperateResult struct{}"
	}

	return strings.Join([]string{"VmOperateResult", string(data)}, " ")
}
