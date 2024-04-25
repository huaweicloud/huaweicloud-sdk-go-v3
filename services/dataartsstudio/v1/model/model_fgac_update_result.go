package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FgacUpdateResult struct {

	// 数据连接id
	Id *string `json:"id,omitempty"`

	// 是否更新成功,true表示更新成功,false表示更新失败。
	Status *bool `json:"status,omitempty"`

	// 细粒度认证更新错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o FgacUpdateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FgacUpdateResult struct{}"
	}

	return strings.Join([]string{"FgacUpdateResult", string(data)}, " ")
}
