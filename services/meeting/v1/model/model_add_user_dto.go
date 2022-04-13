package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 用户信息。
type AddUserDto struct {
	// 企业用户姓名。 maxLength：64 minLength：1

	Name string `json:"name"`
	// 企业用户的英文姓名。 maxLength：64 minLength：0

	EnglishName *string `json:"englishName,omitempty"`
	// 企业用户帐号，若携带则以携带为准，否则后台自动生成。账号整系统唯一 帐号只能包含大小写字母、数字、_、-、.、@符号，不能为纯数字和@后面带.号。 maxLength：64 minLength：0

	Account *string `json:"account,omitempty"`
	// 指定第三方账号开户

	ThirdAccount *string `json:"thirdAccount,omitempty"`
	// 手机号，必须加上国家码。 例如中国大陆手机+86xxxxxxx。当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 说明：手机号或者邮箱至少填写一个 maxLength：32 minLength：0

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 企业用户帐号的密码。若携带则以实际携带为准，否则后台默认生成，密码必须满足： 1、8-32位 2、不能和账号的正序和倒序一致 3、至少包含两种字符类型：小写字母、大写字母、数字、特殊字符（` ~ ! @ # $ % ^ & * ( ) - _ = + | [ { } ] ; : \" ,’ < . > / ?）

	Pwd *string `json:"pwd,omitempty"`
	// 邮箱 maxLength：255 minLength：0

	Email *string `json:"email,omitempty"`
	// 部门编号，若不携带则默认根部门 默认值：1 maxLength：32 minLength：0

	DeptCode *string `json:"deptCode,omitempty"`
	// 签名 maxLength：512 minLength：0

	Signature *string `json:"signature,omitempty"`
	// 职位 maxLength：32 minLength：0

	Title *string `json:"title,omitempty"`
	// 备注 maxLength：128 minLength：0

	Desc *string `json:"desc,omitempty"`
	// 用户状态 * 0、正常 * 1、停用 默认值：0

	Status *AddUserDtoStatus `json:"status,omitempty"`

	Function *UserFunctionDto `json:"function,omitempty"`
	// 是否发送邮件和短信通知， - 0 不发送 - 不填或者其他值就发送, 默认发送  约束 - 长度范围为0到32个字符

	SendNotify *string `json:"sendNotify,omitempty"`
	// 通讯录排序等级，序号越低优先级越高。 默认值：10000 maximum：10000 minimum：1

	SortLevel *int32 `json:"sortLevel,omitempty"`
	// 是否隐藏手机号码 默认值：false

	HidePhone *bool `json:"hidePhone,omitempty"`
}

func (o AddUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserDto struct{}"
	}

	return strings.Join([]string{"AddUserDto", string(data)}, " ")
}

type AddUserDtoStatus struct {
	value int32
}

type AddUserDtoStatusEnum struct {
	E_0 AddUserDtoStatus
	E_1 AddUserDtoStatus
}

func GetAddUserDtoStatusEnum() AddUserDtoStatusEnum {
	return AddUserDtoStatusEnum{
		E_0: AddUserDtoStatus{
			value: 0,
		}, E_1: AddUserDtoStatus{
			value: 1,
		},
	}
}

func (c AddUserDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddUserDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
