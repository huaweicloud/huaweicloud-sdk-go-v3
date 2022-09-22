package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询企业用户信息。
type SearchUserResultDto struct {

	// 用户UUID。
	Id *string `json:"id,omitempty"`

	// 华为云会议帐号。
	UserAccount *string `json:"userAccount,omitempty"`

	// 名称。
	Name *string `json:"name,omitempty"`

	// 英文名称。
	EnglishName *string `json:"englishName,omitempty"`

	// 联系电话。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 个人会议ID。
	VmrId *string `json:"vmrId,omitempty"`

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 部门完整名称。
	DeptNamePath *string `json:"deptNamePath,omitempty"`

	// 用户类型。 * 2：普通用户 * 12：智慧屏用户 * 13：IdeaHub用户 * 14：SmartRooms用户
	UserType *int32 `json:"userType,omitempty"`

	// 管理员类型。 - 0：默认（超级）管理员 - 1：普通管理员 - 2：非管理员（即为普通企业成员，UserType是2时有效）
	AdminType *int32 `json:"adminType,omitempty"`

	// 签名。
	Signature *string `json:"signature,omitempty"`

	// 职位。
	Title *string `json:"title,omitempty"`

	// 备注。
	Desc *string `json:"desc,omitempty"`

	// 用户状态。 * 0：正常 * 1：停用
	Status *int32 `json:"status,omitempty"`

	// 通讯录排序等级，序号越低优先级越高。
	SortLevel *int32 `json:"sortLevel,omitempty"`

	// 是否隐藏手机号码。
	HidePhone *bool `json:"hidePhone,omitempty"`

	// 第三方User ID。
	ThirdAccount *string `json:"thirdAccount,omitempty"`

	// 智慧屏帐号。
	VisionAccount *string `json:"visionAccount,omitempty"`

	// 许可证。 * 0：商用 * 1：免费试用
	License *int32 `json:"license,omitempty"`

	// 激活时间，utc时间戳。
	ActiveTime *int64 `json:"activeTime,omitempty"`

	// 激活码到期时间,utc时间戳。
	ActiveCodeExpireTime *int64 `json:"activeCodeExpireTime,omitempty"`

	// 已激活的终端到期时间,utc时间戳。
	ExpireTime *int64 `json:"expireTime,omitempty"`

	// 激活码。
	ActiveCode *string `json:"activeCode,omitempty"`
}

func (o SearchUserResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchUserResultDto struct{}"
	}

	return strings.Join([]string{"SearchUserResultDto", string(data)}, " ")
}
