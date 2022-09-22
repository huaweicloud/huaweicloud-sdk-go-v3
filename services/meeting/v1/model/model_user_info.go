package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户鉴权信息。
type UserInfo struct {

	// 用户UUID。
	UserId *string `json:"userId,omitempty"`

	// 华为云会议帐号。
	UcloginAccount string `json:"ucloginAccount"`

	// 用户关联的SIP号码。
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// 号码对应的HA1。
	NumberHA1 *string `json:"numberHA1,omitempty"`

	// 用户别名。
	Alias1 *string `json:"alias1,omitempty"`

	// 用户归属的企业ID。
	CompanyId *string `json:"companyId,omitempty"`

	// 用户所在企业归属的SP ID。
	SpId *string `json:"spId,omitempty"`

	// 企业域名。
	CompanyDomain *string `json:"companyDomain,omitempty"`

	// 本地鉴权。
	Realm *string `json:"realm,omitempty"`

	// 用户类型。 * 1：SP管理用户 * 2：企业用户 * 3：免费注册用户 * 10：企业设备用户 * 11：匿名用户 * 12：智慧屏用户 * 13：IdeaHub用户 * 14：电子白板（SmartRooms）用户
	UserType *int32 `json:"userType,omitempty"`

	// 管理员类型。 * 0：默认管理员 * 1：普通管理员 * 2：非管理员，即普通企业成员，userType为2时有效
	AdminType *int32 `json:"adminType,omitempty"`

	// 用户名称。
	Name *string `json:"name,omitempty"`

	// 用户英文名称。
	NameEn *string `json:"nameEn,omitempty"`

	// 标识是否绑定手机。
	IsBindPhone *bool `json:"isBindPhone,omitempty"`

	// 标识是否是免费试用用户。
	FreeUser *bool `json:"freeUser,omitempty"`

	// 第三方的用户帐号。
	ThirdAccount *string `json:"thirdAccount,omitempty"`

	// 智慧屏设备ID。
	VisionAccount *string `json:"visionAccount,omitempty"`

	// 头像链接。
	HeadPictureUrl *string `json:"headPictureUrl,omitempty"`

	// 机机密码，用于智慧屏登录。
	Password *string `json:"password,omitempty"`

	// 用户状态。 * 0：正常 * 1：停用
	Status *int32 `json:"status,omitempty"`

	// 付费用户机机帐号，用于智慧屏登录。
	PaidAccount *string `json:"paidAccount,omitempty"`

	// 付费用户机机密码，用于智慧屏登录。
	PaidPassword *string `json:"paidPassword,omitempty"`

	// 标识是否是WeLink用户。
	WeLinkUser *bool `json:"weLinkUser,omitempty"`

	// 应用ID。
	AppId *string `json:"appId,omitempty"`

	// tr069帐号。
	Tr069Account *string `json:"tr069Account,omitempty"`

	// 企业类型。 * 0：旗舰版 * 5：免费版 * 6：标准版
	CorpType *int32 `json:"corpType,omitempty"`

	// 华为云帐号ID。
	CloudUserId *string `json:"cloudUserId,omitempty"`

	// 标识是否是灰度用户。
	GrayUser *bool `json:"grayUser,omitempty"`
}

func (o UserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
