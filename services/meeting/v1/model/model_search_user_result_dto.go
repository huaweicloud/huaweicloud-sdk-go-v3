package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询企业用户信息
type SearchUserResultDto struct {

	// 用户id
	Id *string `json:"id,omitempty" xml:"id"`

	// 账号
	UserAccount *string `json:"userAccount,omitempty" xml:"userAccount"`

	// 姓名
	Name *string `json:"name,omitempty" xml:"name"`

	// 英文名称
	EnglishName *string `json:"englishName,omitempty" xml:"englishName"`

	// 手机
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 手机所属国家
	Country *string `json:"country,omitempty" xml:"country"`

	// 邮箱
	Email *string `json:"email,omitempty" xml:"email"`

	// 个人会议室Id
	VmrId *string `json:"vmrId,omitempty" xml:"vmrId"`

	// 部门编码
	DeptCode *string `json:"deptCode,omitempty" xml:"deptCode"`

	// 部门完整名称
	DeptNamePath *string `json:"deptNamePath,omitempty" xml:"deptNamePath"`

	// 用户类型 - 0：系统管理员账户 - 1：SP管理员账户 - 2：企业成员账户 - 3：注册用户 - 10：硬终端用户
	UserType *int32 `json:"userType,omitempty" xml:"userType"`

	// 管理员类型 - 0：默认（超级）管理员 - 1：普通管理员 - 2：非管理员（即为普通企业成员，UserType是2时有效）
	AdminType *int32 `json:"adminType,omitempty" xml:"adminType"`

	// 签名
	Signature *string `json:"signature,omitempty" xml:"signature"`

	// 职位
	Title *string `json:"title,omitempty" xml:"title"`

	// 备注
	Desc *string `json:"desc,omitempty" xml:"desc"`

	// 用户状态 * 0：正常 * 1：停用
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 通讯录排序等级，序号越低优先级越高
	SortLevel *int32 `json:"sortLevel,omitempty" xml:"sortLevel"`

	// 是否隐藏手机号码
	HidePhone *bool `json:"hidePhone,omitempty" xml:"hidePhone"`

	// 第三方账号自动开户的第三方账号、Ideahub账号的sn等
	ThirdAccount *string `json:"thirdAccount,omitempty" xml:"thirdAccount"`

	// 智慧屏账号
	VisionAccount *string `json:"visionAccount,omitempty" xml:"visionAccount"`

	// 许可证 * 0：商用 * 1：免费试用
	License *int32 `json:"license,omitempty" xml:"license"`

	// 激活时间，utc时间戳
	ActiveTime *int64 `json:"activeTime,omitempty" xml:"activeTime"`

	// 激活码到期时间,utc时间戳
	ActiveCodeExpireTime *int64 `json:"activeCodeExpireTime,omitempty" xml:"activeCodeExpireTime"`

	// 已激活的终端到期时间,utc时间戳
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime"`

	// 激活码
	ActiveCode *string `json:"activeCode,omitempty" xml:"activeCode"`
}

func (o SearchUserResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchUserResultDto struct{}"
	}

	return strings.Join([]string{"SearchUserResultDto", string(data)}, " ")
}
