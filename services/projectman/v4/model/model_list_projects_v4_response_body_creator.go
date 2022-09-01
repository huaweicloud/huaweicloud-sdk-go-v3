package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建者信息
type ListProjectsV4ResponseBodyCreator struct {

	// 创建人numId
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 创建人id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 创建人姓名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 创建人租户id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 创建人租户名称
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 创建人租户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`
}

func (o ListProjectsV4ResponseBodyCreator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsV4ResponseBodyCreator struct{}"
	}

	return strings.Join([]string{"ListProjectsV4ResponseBodyCreator", string(data)}, " ")
}
