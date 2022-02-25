package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 用户信息。
type ModUserDto struct {
	// 企业用户姓名。 maxLength：64 minLength：1

	Name *string `json:"name,omitempty"`
	// 企业用户的英文姓名。 maxLength：64 minLength：0

	EnglishName *string `json:"englishName,omitempty"`
	// 手机号，必须加上国家码。 例如中国大陆手机为“+86xxxxxxxxxxx”。当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 说明：手机号或者邮箱至少填写一个。 maxLength：32 minLength：0

	Phone *string `json:"phone,omitempty"`
	// 若smsNumber为手机号,则需带上手机号所属的国家。 例如国家为中国大陆则country参数取值为chinaPR 国家和国家码的对应关系请参考：https://support.huaweicloud.com/api-meeting/meeting_21_0109.html

	Country *string `json:"country,omitempty"`
	// 邮箱 maxLength：255 minLength：0

	Email *string `json:"email,omitempty"`
	// 虚拟会议室ID，若不携带则后台默认生成。 maxLength：32 minLength：0

	VmrId *string `json:"vmrId,omitempty"`
	// 部门编号，若不携带则默认根部门 默认值：1 maxLength：32 minLength：0

	DeptCode *string `json:"deptCode,omitempty"`
	// 签名 maxLength：512 minLength：0

	Signature *string `json:"signature,omitempty"`
	// 职位 maxLength：32 minLength：0

	Title *string `json:"title,omitempty"`
	// 备注 maxLength：128 minLength：0

	Desc *string `json:"desc,omitempty"`
	// 用户状态 * 0、正常 * 1、停用 默认值：0

	Status *ModUserDtoStatus `json:"status,omitempty"`
	// 通讯录排序等级，序号越低优先级越高。 默认值：10000 maximum：10000 minimum：1

	SortLevel *int32 `json:"sortLevel,omitempty"`
	// 是否隐藏手机号码 默认值：false

	HidePhone *bool `json:"hidePhone,omitempty"`
}

func (o ModUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModUserDto struct{}"
	}

	return strings.Join([]string{"ModUserDto", string(data)}, " ")
}

type ModUserDtoStatus struct {
	value int32
}

type ModUserDtoStatusEnum struct {
	E_0 ModUserDtoStatus
	E_1 ModUserDtoStatus
}

func GetModUserDtoStatusEnum() ModUserDtoStatusEnum {
	return ModUserDtoStatusEnum{
		E_0: ModUserDtoStatus{
			value: 0,
		}, E_1: ModUserDtoStatus{
			value: 1,
		},
	}
}

func (c ModUserDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModUserDtoStatus) UnmarshalJSON(b []byte) error {
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
