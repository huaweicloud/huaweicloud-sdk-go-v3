package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowResourceInstancesCountRequest struct {

	// 资源类型 policy服务策略 ou组织OU account帐号信息 root根
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
	POLICY  ShowResourceInstancesCountRequestResourceType
	OU      ShowResourceInstancesCountRequestResourceType
	ACCOUNT ShowResourceInstancesCountRequestResourceType
	ROOT    ShowResourceInstancesCountRequestResourceType
}

func GetShowResourceInstancesCountRequestResourceTypeEnum() ShowResourceInstancesCountRequestResourceTypeEnum {
	return ShowResourceInstancesCountRequestResourceTypeEnum{
		POLICY: ShowResourceInstancesCountRequestResourceType{
			value: "policy",
		},
		OU: ShowResourceInstancesCountRequestResourceType{
			value: "ou",
		},
		ACCOUNT: ShowResourceInstancesCountRequestResourceType{
			value: "account",
		},
		ROOT: ShowResourceInstancesCountRequestResourceType{
			value: "root",
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
