package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOpitonsPlacement 实例placement，指定服务器策略，不指定默认从空闲服务器随机选择。
type CreateInstanceOpitonsPlacement struct {

	// 指定服务器
	ServerId *string `json:"server_id,omitempty"`
}

func (o CreateInstanceOpitonsPlacement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOpitonsPlacement struct{}"
	}

	return strings.Join([]string{"CreateInstanceOpitonsPlacement", string(data)}, " ")
}
