package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceInstancesCountRequest Request Object
type ShowResourceInstancesCountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源类型。枚举值：organizations:policies（服务策略）、organizations:ous（组织OU）、organizations:accounts（账号信息） 、organizations:roots：（根）。
	ResourceType ShowResourceInstancesCountRequestResourceType `json:"resource_type"`

	Body *ResourceInstanceReqBody `json:"body,omitempty"`
}

func (o ShowResourceInstancesCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInstancesCountRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceInstancesCountRequest", string(data)}, " ")
}

type ShowResourceInstancesCountRequestResourceType struct {
	value string
}

type ShowResourceInstancesCountRequestResourceTypeEnum struct {
	ORGANIZATIONSPOLICIES ShowResourceInstancesCountRequestResourceType
	ORGANIZATIONSOUS      ShowResourceInstancesCountRequestResourceType
	ORGANIZATIONSACCOUNTS ShowResourceInstancesCountRequestResourceType
	ORGANIZATIONSROOTS    ShowResourceInstancesCountRequestResourceType
}

func GetShowResourceInstancesCountRequestResourceTypeEnum() ShowResourceInstancesCountRequestResourceTypeEnum {
	return ShowResourceInstancesCountRequestResourceTypeEnum{
		ORGANIZATIONSPOLICIES: ShowResourceInstancesCountRequestResourceType{
			value: "organizations:policies",
		},
		ORGANIZATIONSOUS: ShowResourceInstancesCountRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: ShowResourceInstancesCountRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSROOTS: ShowResourceInstancesCountRequestResourceType{
			value: "organizations:roots",
		},
	}
}

func (c ShowResourceInstancesCountRequestResourceType) Value() string {
	return c.value
}

func (c ShowResourceInstancesCountRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceInstancesCountRequestResourceType) UnmarshalJSON(b []byte) error {
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
