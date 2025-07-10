package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTerminalsBindingDesktopsResult struct {

	// 需删除的策略ID。
	Id *string `json:"id,omitempty"`

	// 删除操作的结果码。
	DeleteResultCode *string `json:"delete_result_code,omitempty"`

	// 删除操作的结果信息。
	DeleteResultMsg *string `json:"delete_result_msg,omitempty"`
}

func (o DeleteTerminalsBindingDesktopsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTerminalsBindingDesktopsResult struct{}"
	}

	return strings.Join([]string{"DeleteTerminalsBindingDesktopsResult", string(data)}, " ")
}
