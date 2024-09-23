package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTagResourceRequest Request Object
type DeleteTagResourceRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源类型。枚举值：organizations:policies（服务策略）、organizations:ous（组织OU）、organizations:accounts（账号信息） 、organizations:roots：（根）。
	ResourceType DeleteTagResourceRequestResourceType `json:"resource_type"`

	// 根、组织单元、账号或策略的唯一标识符（ID）。
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
	ORGANIZATIONSPOLICIES DeleteTagResourceRequestResourceType
	ORGANIZATIONSOUS      DeleteTagResourceRequestResourceType
	ORGANIZATIONSACCOUNTS DeleteTagResourceRequestResourceType
	ORGANIZATIONSROOTS    DeleteTagResourceRequestResourceType
}

func GetDeleteTagResourceRequestResourceTypeEnum() DeleteTagResourceRequestResourceTypeEnum {
	return DeleteTagResourceRequestResourceTypeEnum{
		ORGANIZATIONSPOLICIES: DeleteTagResourceRequestResourceType{
			value: "organizations:policies",
		},
		ORGANIZATIONSOUS: DeleteTagResourceRequestResourceType{
			value: "organizations:ous",
		},
		ORGANIZATIONSACCOUNTS: DeleteTagResourceRequestResourceType{
			value: "organizations:accounts",
		},
		ORGANIZATIONSROOTS: DeleteTagResourceRequestResourceType{
			value: "organizations:roots",
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
