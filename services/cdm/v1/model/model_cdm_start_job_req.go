package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmStartJobReq struct {

	// 启动作业，配置变量参数，作业配置无变量时，为空对象
	Variables *interface{} `json:"variables,omitempty"`
}

func (o CdmStartJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmStartJobReq struct{}"
	}

	return strings.Join([]string{"CdmStartJobReq", string(data)}, " ")
}
