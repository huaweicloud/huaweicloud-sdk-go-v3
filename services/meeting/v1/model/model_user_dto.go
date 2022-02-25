package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询企业用户信息
type UserDto struct {
	// 用户ID。

	Id *string `json:"id,omitempty"`
	// 查询用户详情时, 根据不同情况，响应不同。 * 0： 查询成功且用户信息有变化， 响应会把新的信息都返回回去 * 1 ：查询成功且用户信息没有变化，响应只会返回用户ID * 2 ：用户不存在 * 3 ：无权限查询这个用户

	StatusCode *int32 `json:"statusCode,omitempty"`
	// 用户账号。

	Account *string `json:"account,omitempty"`
	// 用户名。

	Name *string `json:"name,omitempty"`
	// 英文名。

	EnglishName *string `json:"englishName,omitempty"`
	// 邮箱。

	Email *string `json:"email,omitempty"`
	// 用户手机。

	Phone *string `json:"phone,omitempty"`
	// 用户部门。

	DeptName *string `json:"deptName,omitempty"`
	// 用户号码。

	Number *string `json:"number,omitempty"`
	// 用户信息最后更新时间。

	UpdateTime *int64 `json:"updateTime,omitempty"`
	// 是否为硬终端。

	IsHardTerminal *bool `json:"isHardTerminal,omitempty"`
	// 用户虚拟会议室ID。

	VmrId *string `json:"vmrId,omitempty"`
	// 用户签名。

	Signature *string `json:"signature,omitempty"`
	// 职位。

	Title *string `json:"title,omitempty"`
	// 描述信息。

	Description *string `json:"description,omitempty"`
	// 是否隐藏手机号（如果为true，其他人查询该用户时，不会返回该用户的手机号。自己查自己是可见的）

	HidePhone *bool `json:"hidePhone,omitempty"`
	// 类型： * NORMAL_USER=普通用户 * HARD_TERMINAL=硬终端用户 * WHITE_BOARD=第三方白板 * HW_VISION_MEMBER=智慧屏

	Type *string `json:"type,omitempty"`
}

func (o UserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDto struct{}"
	}

	return strings.Join([]string{"UserDto", string(data)}, " ")
}
