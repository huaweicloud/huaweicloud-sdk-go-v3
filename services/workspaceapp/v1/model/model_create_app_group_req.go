package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupReq 创建应用组,允许创建应用组之后再绑定应用服务器组。
type CreateAppGroupReq struct {

	// 应用组名称,名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成。 2. 长度范围1~64个字符。
	Name string `json:"name"`

	// 应用服务器组ID。
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 应用组描述。
	Description *string `json:"description,omitempty"`

	AuthorizationType *AuthorizationTypeEnum `json:"authorization_type,omitempty"`

	AppType *AppTypeEnum `json:"app_type,omitempty"`
}

func (o CreateAppGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupReq struct{}"
	}

	return strings.Join([]string{"CreateAppGroupReq", string(data)}, " ")
}
