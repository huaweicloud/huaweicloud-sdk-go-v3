package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTagResourceRequest Request Object
type CreateTagResourceRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 资源类型 identitycenter:permissionset权限集
	ResourceType CreateTagResourceRequestResourceType `json:"resource_type"`

	// 权限集的唯一标识符（ID）。
	ResourceId string `json:"resource_id"`

	Body *TagResourceReqBody `json:"body,omitempty"`
}

func (o CreateTagResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateTagResourceRequest", string(data)}, " ")
}

type CreateTagResourceRequestResourceType struct {
	value string
}

type CreateTagResourceRequestResourceTypeEnum struct {
	IDENTITYCENTERPERMISSIONSET CreateTagResourceRequestResourceType
}

func GetCreateTagResourceRequestResourceTypeEnum() CreateTagResourceRequestResourceTypeEnum {
	return CreateTagResourceRequestResourceTypeEnum{
		IDENTITYCENTERPERMISSIONSET: CreateTagResourceRequestResourceType{
			value: "identitycenter:permissionset",
		},
	}
}

func (c CreateTagResourceRequestResourceType) Value() string {
	return c.value
}

func (c CreateTagResourceRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagResourceRequestResourceType) UnmarshalJSON(b []byte) error {
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
