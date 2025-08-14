package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountAssignmentsForPrincipalRequest Request Object
type ListAccountAssignmentsForPrincipalRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
	InstanceId string `json:"instance_id"`

	// IAM Identity Center 身份主体的全局唯一标识符（ID）
	PrincipalId string `json:"principal_id"`

	// IAM Identity Center 身份主体类型.
	PrincipalType ListAccountAssignmentsForPrincipalRequestPrincipalType `json:"principal_type"`

	// 每个请求返回的最大结果数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 指定账户的唯一身份标识
	AccountId *string `json:"account_id,omitempty"`
}

func (o ListAccountAssignmentsForPrincipalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentsForPrincipalRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentsForPrincipalRequest", string(data)}, " ")
}

type ListAccountAssignmentsForPrincipalRequestPrincipalType struct {
	value string
}

type ListAccountAssignmentsForPrincipalRequestPrincipalTypeEnum struct {
	GROUP ListAccountAssignmentsForPrincipalRequestPrincipalType
	USER  ListAccountAssignmentsForPrincipalRequestPrincipalType
}

func GetListAccountAssignmentsForPrincipalRequestPrincipalTypeEnum() ListAccountAssignmentsForPrincipalRequestPrincipalTypeEnum {
	return ListAccountAssignmentsForPrincipalRequestPrincipalTypeEnum{
		GROUP: ListAccountAssignmentsForPrincipalRequestPrincipalType{
			value: "GROUP",
		},
		USER: ListAccountAssignmentsForPrincipalRequestPrincipalType{
			value: "USER",
		},
	}
}

func (c ListAccountAssignmentsForPrincipalRequestPrincipalType) Value() string {
	return c.value
}

func (c ListAccountAssignmentsForPrincipalRequestPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccountAssignmentsForPrincipalRequestPrincipalType) UnmarshalJSON(b []byte) error {
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
