package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddUserResponse struct {
	// 用户id

	Id *string `json:"id,omitempty"`
	// 账号

	UserAccount *string `json:"userAccount,omitempty"`
	// 姓名

	Name *string `json:"name,omitempty"`
	// 英文名称

	EnglishName *string `json:"englishName,omitempty"`
	// 联系电话

	Phone *string `json:"phone,omitempty"`
	// 联系电话所属的国家

	Country *string `json:"country,omitempty"`
	// 邮箱

	Email *string `json:"email,omitempty"`
	// SIP号码

	SipNum *string `json:"sipNum,omitempty"`
	// 云会议室列表

	VmrList *[]UserVmrDto `json:"vmrList,omitempty"`
	// 部门编码

	DeptCode *string `json:"deptCode,omitempty"`
	// 部门名称

	DeptName *string `json:"deptName,omitempty"`
	// 部门完整名称

	DeptNamePath *string `json:"deptNamePath,omitempty"`
	// 用户类型 - 2：企业成员账户

	UserType *int32 `json:"userType,omitempty"`
	// 管理员类型 - 0：默认（超级）管理员 - 1：普通管理员 - 2：非管理员（即为普通企业成员，UserType是2时有效）

	AdminType *int32 `json:"adminType,omitempty"`
	// 签名

	Signature *string `json:"signature,omitempty"`
	// 职位

	Title *string `json:"title,omitempty"`
	// 备注

	Desc *string `json:"desc,omitempty"`

	Corp *CorpBasicInfoDto `json:"corp,omitempty"`

	Function *UserFunctionDto `json:"function,omitempty"`

	DevType *QueryDeviceInfoResultDto `json:"devType,omitempty"`
	// 用户状态 * 0、正常 * 1、停用

	Status *int32 `json:"status,omitempty"`
	// 通讯录排序等级，序号越低优先级越高

	SortLevel *int32 `json:"sortLevel,omitempty"`
	// 是否隐藏手机号码

	HidePhone *bool `json:"hidePhone,omitempty"`
	// 智慧屏唯一账号

	VisionAccount *string `json:"visionAccount,omitempty"`
	// 第三方账号，自动开户的第三方账号、Ideahub账号的sn等

	ThirdAccount *string `json:"thirdAccount,omitempty"`
	// 许可证 * 0：商用； * 1：免费试用。

	License *int32 `json:"license,omitempty"`
	// 激活时间，utc时间戳

	ActiveTime     *int64 `json:"activeTime,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserResponse struct{}"
	}

	return strings.Join([]string{"AddUserResponse", string(data)}, " ")
}
