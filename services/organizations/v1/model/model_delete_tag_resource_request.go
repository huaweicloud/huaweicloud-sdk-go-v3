package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteTagResourceRequest struct {

	// 资源类型 organizations:policies服务策略 organizations:ous组织OU organizations:accounts 帐号信息 organizations:roots根
	ResourceType DeleteTagResourceRequestResourceType `json:"resource_type"`

	// 根、组织单元、帐号或策略的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o DeleteTagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagResourceRequest", string(data)}, " ")
}

type DeleteTagResourceRequestResourceType struct {
	value string
}

type DeleteTagResourceRequestResourceTypeEnum struct {
	ORGANIZATIONSROOTS    DeleteTagResourceRequestResourceType
	ORGANIZATIONSOUS      DeleteTagResourceRequestResourceType
	ORGANIZATIONSACCOUNTS DeleteTagResourceRequestResourceType
	ORGANIZATIONSPOLICIES DeleteTagResourceRequestResourceType
}

func GetDeleteTagResourceRequestResourceTypeEnum() DeleteTagResourceRequestResourceTypeEnum {
	return DeleteTagResourceRequestResourceTypeEnum{
		ORGANIZATIONSROOTS: DeleteTagResourceRequestResourceType{
			value: "organizations:roots",
		},
		ORGANIZATIONSOUS: DeleteTagResourceRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: DeleteTagResourceRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSPOLICIES: DeleteTagResourceRequestResourceType{
			value: "organizations:policies",
		},
	}
}

func (c DeleteTagResourceRequestResourceType) Value() string {
	return c.value
}

func (c DeleteTagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
