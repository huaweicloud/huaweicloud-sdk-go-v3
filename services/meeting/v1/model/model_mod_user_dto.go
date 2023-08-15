package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModUserDto 用户信息。
type ModUserDto struct {

	// 企业用户名称。
	Name *string `json:"name,omitempty"`

	// 企业用户的英文名称。
	EnglishName *string `json:"englishName,omitempty"`

	// 手机号，必须加上国家码。 例如中国大陆手机为“+86xxxxxxxxxxx”。当填写手机号时 “country”参数必填。 手机号只允许输入纯数字。 说明：手机号或者邮箱至少填写一个。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 邮箱地址。
	Email *string `json:"email,omitempty"`

	// 个人会议ID，若不携带则后台默认生成。
	VmrId *string `json:"vmrId,omitempty"`

	// 个人会议ID，若不携带则后台默认生成。 默认值：1
	DeptCode *string `json:"deptCode,omitempty"`

	// 签名。
	Signature *string `json:"signature,omitempty"`

	// 职位。
	Title *string `json:"title,omitempty"`

	// 备注。
	Desc *string `json:"desc,omitempty"`

	// 用户状态。默认值：0。 * 0：正常 * 1：停用
	Status *ModUserDtoStatus `json:"status,omitempty"`

	// 通讯录排序等级，序号越低优先级越高。 默认值：10000
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

func (c ModUserDtoStatus) Value() int32 {
	return c.value
}

func (c ModUserDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModUserDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
