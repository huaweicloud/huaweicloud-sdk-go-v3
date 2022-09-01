package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户鉴权信息。
type UserInfo struct {

	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId"`

	// 用户UC帐号
	UcloginAccount string `json:"ucloginAccount" xml:"ucloginAccount"`

	// 用户关联的号码，sip格式。 登录类型不一样获取到的号码也不同，如软终端和硬终端客户端登录获取的号码不同。若未关联号码，则为空。 例如：sip:+8675590121000@domain5.huawei.com
	ServiceAccount *string `json:"serviceAccount,omitempty" xml:"serviceAccount"`

	// 号码对应的HA1
	NumberHA1 *string `json:"numberHA1,omitempty" xml:"numberHA1"`

	// 用户别名1
	Alias1 *string `json:"alias1,omitempty" xml:"alias1"`

	// 企业ID
	CompanyId *string `json:"companyId,omitempty" xml:"companyId"`

	// SP ID
	SpId *string `json:"spId,omitempty" xml:"spId"`

	// 企业域名
	CompanyDomain *string `json:"companyDomain,omitempty" xml:"companyDomain"`

	// 本地鉴权：realm
	Realm *string `json:"realm,omitempty" xml:"realm"`

	// 用户类型。 * 0：系统管理用户 * 1：SP管理用户 * 2：企业用户 * 3：upath用户 * 4：硬终端默认用户 * 5：TE终端用户 * 6：顾客用户 * 7：公共设备用户 * 8：集群群组用户 * 9：USM用户
	UserType *int32 `json:"userType,omitempty" xml:"userType"`

	// 管理员类型： * 0：默认管理员 * 1：普通管理员 * 2：非管理员，即普通企业成员，USERTYPE为2时有效
	AdminType *int32 `json:"adminType,omitempty" xml:"adminType"`

	// 用户姓名
	Name *string `json:"name,omitempty" xml:"name"`

	// 用户英文姓名
	NameEn *string `json:"nameEn,omitempty" xml:"nameEn"`

	// 标识是否绑定手机
	IsBindPhone *bool `json:"isBindPhone,omitempty" xml:"isBindPhone"`

	// 标识是否是免费试用用户
	FreeUser *bool `json:"freeUser,omitempty" xml:"freeUser"`

	// 用户的第三方账号，例如华为账号登录时获取到的union_id
	ThirdAccount *string `json:"thirdAccount,omitempty" xml:"thirdAccount"`

	// 智慧屏设备id
	VisionAccount *string `json:"visionAccount,omitempty" xml:"visionAccount"`

	// 头像链接
	HeadPictureUrl *string `json:"headPictureUrl,omitempty" xml:"headPictureUrl"`

	// 机机密码，用于智慧屏登录
	Password *string `json:"password,omitempty" xml:"password"`

	// 用户状态。 * 0：正常 * 1：停用
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 付费用户机机账号，用于智慧屏登录
	PaidAccount *string `json:"paidAccount,omitempty" xml:"paidAccount"`

	// 付费用户机机密码，用于智慧屏登录
	PaidPassword *string `json:"paidPassword,omitempty" xml:"paidPassword"`

	// 标识是否是WeLink用户
	WeLinkUser *bool `json:"weLinkUser,omitempty" xml:"weLinkUser"`

	// 应用ID
	AppId *string `json:"appId,omitempty" xml:"appId"`

	// tr069帐号
	Tr069Account *string `json:"tr069Account,omitempty" xml:"tr069Account"`

	// 企业类型。 * 0：企业版 * 1：公共企业，手机、邮箱注册时会放到该企业内 * 2：公共企业，智慧屏用户自动开户时会放到该企业内 * 3：公共企业，大屏用户自动开户时会放到该企业内 * 4：公共TOC消费者企业 * 5：免费版 * 6：专业版
	CorpType *int32 `json:"corpType,omitempty" xml:"corpType"`

	// 华为云账号ID
	CloudUserId *string `json:"cloudUserId,omitempty" xml:"cloudUserId"`

	// 标识是否是灰度用户
	GrayUser *bool `json:"grayUser,omitempty" xml:"grayUser"`
}

func (o UserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
