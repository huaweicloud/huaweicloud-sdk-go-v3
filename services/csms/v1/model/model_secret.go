package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Secret 凭据对象。
type Secret struct {

	// 凭据的资源标识符。
	Id *string `json:"id,omitempty"`

	// 凭据名称。
	Name *string `json:"name,omitempty"`

	// 凭据状态，取值如下：  ENABLED：表示启用状态  DISABLED：表示禁用状态  PENDING_DELETE：表示待删除状态  FROZEN：表示冻结状态
	State *string `json:"state,omitempty"`

	// 用于加密凭据值的KMS主密钥的ID值。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 凭据的描述信息。
	Description *string `json:"description,omitempty"`

	// 凭据创建时间，时间戳，即从1970年1月1日至该时间的总秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 凭据上次更新时间，时间戳，即从1970年1月1日至该时间的总秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 凭据计划删除时间，时间戳，即从1970年1月1日至该时间的总秒数。  凭据不在删除计划中时，本项值为null。
	ScheduledDeleteTime *int64 `json:"scheduled_delete_time,omitempty"`

	// 凭据类型  - COMMON：通用凭据(默认)。用于应用系统中的各种敏感信息储存。 - RDS：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。（已不支持，使用RDS-FG替代） - RDS-FG：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。 - GaussDB-FG：TaurusDB凭据。专门针对TaurusDB的凭据，用于存储TaurusDB的账号信息。
	SecretType *SecretSecretType `json:"secret_type,omitempty"`

	// 自动轮转  取值：true 开启, false 关闭(默认)
	AutoRotation *bool `json:"auto_rotation,omitempty"`

	// 轮转周期  约束：6小时-8,760小时 （365天）  类型：Integer[unit] ，Integer表示时间长度 。unit表示时间单位，d（天）、h（小时）、m（分钟）、s（秒）。例如 1d 表示一天，24h也表示一天  说明：当开启自动轮转时，必须填写该值
	RotationPeriod *string `json:"rotation_period,omitempty"`

	// 轮转配置  约束：范围不超过1024个字符。  当secret_type为RDS-FG、GaussDB-FG时，配置为{\"InstanceId\":\"\",\"SecretSubType\":\"\"}  说明：当secret_type为RDS-FG、GaussDB-FG时，必须填写该值  InstanceId为实例ID,SecretSubType为轮转子类型，取值为：SingleUser，MultiUser。  SingleUser：指定轮转类型为单用户模式轮转，每次轮转将指定账号重置为新的口令。  MultiUser：指定轮转类型为双用户模式轮转，SYSCURRENT和SYSPREVIOUS分别引用其中一个账号。凭据轮转时，SYSPREVIOUS引用的账号口令会被重置为新的随机口令，随后凭据交换SYSCURRENT和SYSPREVIOUS对账号的引用。
	RotationConfig *string `json:"rotation_config,omitempty"`

	// 轮转时间戳
	RotationTime *int64 `json:"rotation_time,omitempty"`

	// 下一次轮转时间戳
	NextRotationTime *int64 `json:"next_rotation_time,omitempty"`

	// 凭据订阅的事件列表，当前最大可订阅一个事件。当事件包含的基础事件触发时，通知消息将发送到事件对应的通知主题。
	EventSubscriptions *[]string `json:"event_subscriptions,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// FunctionGraph函数的urn。
	RotationFuncUrn *string `json:"rotation_func_urn,omitempty"`
}

func (o Secret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secret struct{}"
	}

	return strings.Join([]string{"Secret", string(data)}, " ")
}

type SecretSecretType struct {
	value string
}

type SecretSecretTypeEnum struct {
	COMMON      SecretSecretType
	RDS_FG      SecretSecretType
	GAUSS_DB_FG SecretSecretType
}

func GetSecretSecretTypeEnum() SecretSecretTypeEnum {
	return SecretSecretTypeEnum{
		COMMON: SecretSecretType{
			value: "COMMON",
		},
		RDS_FG: SecretSecretType{
			value: "RDS-FG",
		},
		GAUSS_DB_FG: SecretSecretType{
			value: "GaussDB-FG",
		},
	}
}

func (c SecretSecretType) Value() string {
	return c.value
}

func (c SecretSecretType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecretSecretType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
