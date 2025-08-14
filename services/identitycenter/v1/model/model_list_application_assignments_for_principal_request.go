package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApplicationAssignmentsForPrincipalRequest Request Object
type ListApplicationAssignmentsForPrincipalRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// 身份主体的全局唯一标识符（ID）
	PrincipalId string `json:"principal_id"`

	PrincipalType ListApplicationAssignmentsForPrincipalRequestPrincipalType `json:"principal_type"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListApplicationAssignmentsForPrincipalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAssignmentsForPrincipalRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationAssignmentsForPrincipalRequest", string(data)}, " ")
}

type ListApplicationAssignmentsForPrincipalRequestPrincipalType struct {
	value string
}

type ListApplicationAssignmentsForPrincipalRequestPrincipalTypeEnum struct {
	USER  ListApplicationAssignmentsForPrincipalRequestPrincipalType
	GROUP ListApplicationAssignmentsForPrincipalRequestPrincipalType
}

func GetListApplicationAssignmentsForPrincipalRequestPrincipalTypeEnum() ListApplicationAssignmentsForPrincipalRequestPrincipalTypeEnum {
	return ListApplicationAssignmentsForPrincipalRequestPrincipalTypeEnum{
		USER: ListApplicationAssignmentsForPrincipalRequestPrincipalType{
			value: "USER",
		},
		GROUP: ListApplicationAssignmentsForPrincipalRequestPrincipalType{
			value: "GROUP",
		},
	}
}

func (c ListApplicationAssignmentsForPrincipalRequestPrincipalType) Value() string {
	return c.value
}

func (c ListApplicationAssignmentsForPrincipalRequestPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApplicationAssignmentsForPrincipalRequestPrincipalType) UnmarshalJSON(b []byte) error {
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
