package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListKeystoreSearchResponseBodyResultKeystoreList struct {

	// 文件ID
	Id *string `json:"id,omitempty"`

	Permission *ListKeystoreSearchResponseBodyResultPermission `json:"permission,omitempty"`

	// 创建时间戳
	CreateTimeStamp *string `json:"create_time_stamp,omitempty"`

	// 修改时间戳
	UpdateTimeStamp *string `json:"update_time_stamp,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`

	// 文件名
	KeystoreName *string `json:"keystore_name,omitempty"`

	// 文件创建者用户ID
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 文件创建者用户名
	CreateUserName *string `json:"create_user_name,omitempty"`

	// 文件修改者用户ID
	UpdateUserId *string `json:"update_user_id,omitempty"`

	// 文件修改者用户名
	UpdateUserName *string `json:"update_user_name,omitempty"`

	// 是否共享，开启后允许租户内所有成员在编译构建中使用该文件
	Share float32 `json:"share,omitempty"`

	// 更新时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o ListKeystoreSearchResponseBodyResultKeystoreList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreSearchResponseBodyResultKeystoreList struct{}"
	}

	return strings.Join([]string{"ListKeystoreSearchResponseBodyResultKeystoreList", string(data)}, " ")
}
