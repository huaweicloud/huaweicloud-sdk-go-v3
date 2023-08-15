package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppGroupReq 更新应用
type UpdateAppGroupReq struct {

	// 应用组名称,名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成 2. 长度范围1~64个字符
	Name *string `json:"name,omitempty"`

	// 应用服务器组ID(仅允许未设置的情形下进行绑定)
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 应用组描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateAppGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppGroupReq struct{}"
	}

	return strings.Join([]string{"UpdateAppGroupReq", string(data)}, " ")
}
