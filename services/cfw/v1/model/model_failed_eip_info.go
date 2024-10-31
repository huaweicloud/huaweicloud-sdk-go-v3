package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedEipInfo struct {

	// 修改状态失败的eipId。
	Id *string `json:"id,omitempty"`

	// 修改状态失败的错误码。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o FailedEipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedEipInfo struct{}"
	}

	return strings.Join([]string{"FailedEipInfo", string(data)}, " ")
}
