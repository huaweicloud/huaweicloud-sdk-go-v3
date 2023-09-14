package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretRequestBody 创建凭据请求消息体。
type CreateSecretRequestBody struct {

	// 待创建凭据的名称。  约束：取值范围为1到64个字符，满足正则匹配“^[a-zA-Z0-9_-]{1,64}$”。
	Name string `json:"name"`

	// 用于加密保护凭据值的KMS主密钥ID，如果您未指定此参数，凭据管理服务将默认使用名为csms/default的默认主密钥，用于加密您账号在本项目中创建的凭据值。如果用户账号下不存在该名称的主密钥，则凭据管理服务自动为您创建该名称的密钥。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 凭据的描述信息。  约束：2048字节。
	Description *string `json:"description,omitempty"`

	// 二进制类型凭据在base64编码后的明文，凭据管理服务将其加密后，存入凭据的初始版本中。  类型：base64编码的二进制数据对象。  约束：secret_binary和secret_string必须且只能设置一个，最大32K。 当secret_type为RDS时。凭据值格式为： \"{'users':[{'name':'','password':''}]}\" 其中name为RDS实例账号名称，password为RDS实例账号口令
	SecretBinary *string `json:"secret_binary,omitempty"`

	// 文本类型凭据的明文，凭据管理服务将其加密后，存入凭据的初始版本中。  约束：secret_binary和secret_string必须且只能设置一个，最大32K。
	SecretString *string `json:"secret_string,omitempty"`

	// 凭据类型  取值 ： COMMON ：通用凭据(默认)。用于应用系统中的各种敏感信息储存。         RDS ：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。
	SecretType *string `json:"secret_type,omitempty"`

	// 自动轮转  取值：true 开启 ,false 关闭 (默认)
	AutoRotation *bool `json:"auto_rotation,omitempty"`

	// 轮转周期  约束：6小时-8,760小时 （365天）  类型：Integer[unit] ，Integer表示时间长度 。unit表示时间单位，d（天）、h（小时）、m（分钟）、s（秒）。例如 1d 表示一天，24h也表示一天  说明：当开启自动轮转时，必须填写该值
	RotationPeriod *string `json:"rotation_period,omitempty"`

	// 轮转配置  约束：范围不超过1024个字符。  当secret_type为RDS时，配置为{\"rds_instance_id\":\"\",\"Secret_sub_type\":\"\"}  说明：当secret_type为RDS时，必须填写该值  rds_instance_id为RDS的实例ID,Secret_sub_type为轮转子类型，取值为：SingleUser，MultiUser。  SingleUser：指定轮转类型为单用户模式轮转，每次轮转将指定账号重置为新的口令。  MultiUser：指定轮转类型为双用户模式轮转，SYSCURRENT和SYSPREVIOUS分别引用其中一个账号。凭据轮转时，SYSPREVIOUS引用的账号口令会被重置为新的随机口令，随后凭据交换SYSCURRENT和SYSPREVIOUS对RDS账号的引用。
	RotationConfig *string `json:"rotation_config,omitempty"`

	// 凭据订阅的事件列表，当前最大可订阅一个事件。当事件包含的基础事件触发时，通知消息将发送到事件对应的通知主题。
	EventSubscriptions *[]string `json:"event_subscriptions,omitempty"`

	// 该参数针对企业用户使用。如果您是企业用户，且已创建企业项目，则请从下拉列表中为密钥选择需要绑定的企业项目，默认项目为“default”。 未开通企业管理的用户页面则没有“企业项目”参数项，无需进行配置。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateSecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecretRequestBody", string(data)}, " ")
}
