package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowNamespaceResponse struct {
	// id

	Id *int32 `json:"id,omitempty"`
	// 组织名称

	Name *string `json:"name,omitempty"`
	// IAM用户名

	CreatorName *string `json:"creator_name,omitempty"`
	// 用户权限。7表示管理权限，3表示编辑权限，1表示读取权限。

	Auth           *ShowNamespaceResponseAuth `json:"auth,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowNamespaceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNamespaceResponse struct{}"
	}

	return strings.Join([]string{"ShowNamespaceResponse", string(data)}, " ")
}

type ShowNamespaceResponseAuth struct {
	value int32
}

type ShowNamespaceResponseAuthEnum struct {
	E_7 ShowNamespaceResponseAuth
	E_3 ShowNamespaceResponseAuth
	E_1 ShowNamespaceResponseAuth
}

func GetShowNamespaceResponseAuthEnum() ShowNamespaceResponseAuthEnum {
	return ShowNamespaceResponseAuthEnum{
		E_7: ShowNamespaceResponseAuth{
			value: 7,
		}, E_3: ShowNamespaceResponseAuth{
			value: 3,
		}, E_1: ShowNamespaceResponseAuth{
			value: 1,
		},
	}
}

func (c ShowNamespaceResponseAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowNamespaceResponseAuth) UnmarshalJSON(b []byte) error {
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
