package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 密钥对详细信息
type KeypairDetail struct {

	// SSH密钥对的名称
	Name *string `json:"name,omitempty" xml:"name"`

	// SSH密钥对的ID
	Id *int64 `json:"id,omitempty" xml:"id"`

	// SSH密钥对的类型
	Type *KeypairDetailType `json:"type,omitempty" xml:"type"`

	// 租户级或者用户级
	Scope *KeypairDetailScope `json:"scope,omitempty" xml:"scope"`

	// SSH密钥对对应的publicKey信息
	PublicKey *string `json:"public_key,omitempty" xml:"public_key"`

	// SSH密钥对应指纹信息
	Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint"`

	// 是否托管密钥
	IsKeyProtection *bool `json:"is_key_protection,omitempty" xml:"is_key_protection"`

	// SSH密钥对删除的标记
	Deleted *bool `json:"deleted,omitempty" xml:"deleted"`

	// SSH密钥对的描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// SSH密钥对所属的用户信息
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// SSH密钥对创建的时间，时间戳，即从1970年1月1日至该时间的总秒数
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// SSH密钥对删除的时间，时间戳，即从1970年1月1日至该时间的总秒数
	DeleteTime *int64 `json:"delete_time,omitempty" xml:"delete_time"`

	// SSH密钥对的更新时间，时间戳，即从1970年1月1日至该时间的总秒数
	UpdateTime *int64 `json:"update_time,omitempty" xml:"update_time"`

	// 冻结状态 - 0：正常状态 - 1：普通冻结 - 2：公安冻结 - 3：普通冻结及公安冻结 - 4：违规冻结 - 5：普通冻结及违规冻结 - 6：公安冻结及违规冻结 - 7：普通冻结、公安冻结及违规冻结 - 8：未实名认证冻结 - 9：普通冻结及未实名认证冻结 - 10：公安冻结及未实名认证冻结
	FrozenState *int32 `json:"frozen_state,omitempty" xml:"frozen_state"`
}

func (o KeypairDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeypairDetail struct{}"
	}

	return strings.Join([]string{"KeypairDetail", string(data)}, " ")
}

type KeypairDetailType struct {
	value string
}

type KeypairDetailTypeEnum struct {
	SSH  KeypairDetailType
	X509 KeypairDetailType
}

func GetKeypairDetailTypeEnum() KeypairDetailTypeEnum {
	return KeypairDetailTypeEnum{
		SSH: KeypairDetailType{
			value: "ssh",
		},
		X509: KeypairDetailType{
			value: "x509",
		},
	}
}

func (c KeypairDetailType) Value() string {
	return c.value
}

func (c KeypairDetailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeypairDetailType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type KeypairDetailScope struct {
	value string
}

type KeypairDetailScopeEnum struct {
	DOMAIN KeypairDetailScope
	USER   KeypairDetailScope
}

func GetKeypairDetailScopeEnum() KeypairDetailScopeEnum {
	return KeypairDetailScopeEnum{
		DOMAIN: KeypairDetailScope{
			value: "domain",
		},
		USER: KeypairDetailScope{
			value: "user",
		},
	}
}

func (c KeypairDetailScope) Value() string {
	return c.value
}

func (c KeypairDetailScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeypairDetailScope) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
