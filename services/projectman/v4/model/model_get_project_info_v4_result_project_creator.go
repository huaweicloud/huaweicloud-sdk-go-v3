package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetProjectInfoV4ResultProjectCreator 创建者信息
type GetProjectInfoV4ResultProjectCreator struct {

	// 创建人num_id
	UserNumId *int32 `json:"user_num_id,omitempty"`

	// 创建人uuid
	UserId *string `json:"user_id,omitempty"`

	// 创建人姓名
	UserName *string `json:"user_name,omitempty"`

	// 创建人租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 创建人租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 创建人昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o GetProjectInfoV4ResultProjectCreator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetProjectInfoV4ResultProjectCreator struct{}"
	}

	return strings.Join([]string{"GetProjectInfoV4ResultProjectCreator", string(data)}, " ")
}
