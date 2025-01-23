package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPortRequest struct {

	// 端口号。  TaurusDB端口号范围：大于等于1024，小于等于65535，不包含端口5342-5345、12017、20000、20201、20202、33062。
	Port int32 `json:"port"`
}

func (o ModifyPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPortRequest struct{}"
	}

	return strings.Join([]string{"ModifyPortRequest", string(data)}, " ")
}
