package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailedToBindEipChildInfo 绑定EIP失败的子任务信息
type FailedToBindEipChildInfo struct {

	// 子任务ID
	Id *string `json:"id,omitempty"`

	// 子任务名称
	Name *string `json:"name,omitempty"`
}

func (o FailedToBindEipChildInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedToBindEipChildInfo struct{}"
	}

	return strings.Join([]string{"FailedToBindEipChildInfo", string(data)}, " ")
}
