package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护组故障切换请求体
type FailoverProtectionGroupRequestBody struct {

	// 标识保护组故障切换操作。该参数目前默认值为空。
	FailoverServerGroup *interface{} `json:"failover-server-group" xml:"failover-server-group"`
}

func (o FailoverProtectionGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"FailoverProtectionGroupRequestBody", string(data)}, " ")
}
