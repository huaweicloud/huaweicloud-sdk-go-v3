package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询企业用户信息
type UserDto struct {

	// 用户ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 查询用户详情时, 根据不同情况，响应不同。 * 0： 查询成功且用户信息有变化， 响应会把新的信息都返回回去 * 1 ：查询成功且用户信息没有变化，响应只会返回用户ID * 2 ：用户不存在 * 3 ：无权限查询这个用户
	StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode"`

	// 用户账号。
	Account *string `json:"account,omitempty" xml:"account"`

	// 用户名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 英文名。
	EnglishName *string `json:"englishName,omitempty" xml:"englishName"`

	// 邮箱。
	Email *string `json:"email,omitempty" xml:"email"`

	// 用户手机。
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 用户部门。
	DeptName *string `json:"deptName,omitempty" xml:"deptName"`

	// 用户号码。
	Number *string `json:"number,omitempty" xml:"number"`

	// 用户信息最后更新时间。
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime"`

	// 是否为硬终端。
	IsHardTerminal *bool `json:"isHardTerminal,omitempty" xml:"isHardTerminal"`

	// 用户虚拟会议室ID。
	VmrId *string `json:"vmrId,omitempty" xml:"vmrId"`

	// 用户签名。
	Signature *string `json:"signature,omitempty" xml:"signature"`

	// 职位。
	Title *string `json:"title,omitempty" xml:"title"`

	// 描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否隐藏手机号（如果为true，其他人查询该用户时，不会返回该用户的手机号。自己查自己是可见的）
	HidePhone *bool `json:"hidePhone,omitempty" xml:"hidePhone"`

	// 类型： * NORMAL_USER=普通用户 * HARD_TERMINAL=硬终端用户 * WHITE_BOARD=第三方白板 * HW_VISION_MEMBER=智慧屏
	Type *string `json:"type,omitempty" xml:"type"`

	// 部门编码列表
	DeptCodes *[]string `json:"deptCodes,omitempty" xml:"deptCodes"`
}

func (o UserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDto struct{}"
	}

	return strings.Join([]string{"UserDto", string(data)}, " ")
}
