package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCurUserInfoResponse struct {

	// 租户id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 用户数字id
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time"`

	// 更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty" xml:"updated_time"`

	// 性别
	Gender *string `json:"gender,omitempty" xml:"gender"`

	// 用户类型, User 云用户, Federation 联邦账号,
	UserType       *string `json:"user_type,omitempty" xml:"user_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCurUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCurUserInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCurUserInfoResponse", string(data)}, " ")
}
