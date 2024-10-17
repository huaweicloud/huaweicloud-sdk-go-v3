package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateLogInfo struct {

	// 操作日志ID
	Id *string `json:"id,omitempty"`

	// 操作日志用户名
	User *string `json:"user,omitempty"`

	// 该条记录发生的时间，格式为时间戳。
	Time *string `json:"time,omitempty"`

	// 该条记录的操作类型 - create：创建 - update：更新 - delete：删除 - download: 下载
	Action *string `json:"action,omitempty"`

	// 该条记录的功能类型
	Function *string `json:"function,omitempty"`

	// 该条记录对应的用户操作对象
	Name *string `json:"name,omitempty"`

	// 该条记录具体的描述
	Description *string `json:"description,omitempty"`

	// 该条记录对应用户执行的结果 - success: 成功 - fail: 失败
	Result *string `json:"result,omitempty"`
}

func (o OperateLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateLogInfo struct{}"
	}

	return strings.Join([]string{"OperateLogInfo", string(data)}, " ")
}
