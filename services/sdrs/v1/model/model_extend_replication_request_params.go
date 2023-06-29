package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendReplicationRequestParams 复制对扩容请求数据结构
type ExtendReplicationRequestParams struct {

	// 复制对内的磁盘扩容后的最终容量。单位：GB 说明:该参数的取值为小数时，系统默认取小数点前的整数值。
	NewSize int32 `json:"new_size"`
}

func (o ExtendReplicationRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendReplicationRequestParams struct{}"
	}

	return strings.Join([]string{"ExtendReplicationRequestParams", string(data)}, " ")
}
