package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchResourceShareAssociationsReqBody The request body of the SearchResourceShareAssociations operation.
type SearchResourceShareAssociationsReqBody struct {

	// 指定绑定的状态。
	AssociationStatus *string `json:"association_status,omitempty"`

	// 指定绑定的类型（principal或resource）。
	AssociationType SearchResourceShareAssociationsReqBodyAssociationType `json:"association_type"`

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 指定绑定的资源使用者。
	Principal *string `json:"principal,omitempty"`

	// 指定绑定的共享资源URN。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 指定资源共享实例的ID列表。
	ResourceShareIds *[]string `json:"resource_share_ids,omitempty"`

	// 指定共享资源ID列表。
	ResourceIds *[]string `json:"resource_ids,omitempty"`
}

func (o SearchResourceShareAssociationsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareAssociationsReqBody struct{}"
	}

	return strings.Join([]string{"SearchResourceShareAssociationsReqBody", string(data)}, " ")
}

type SearchResourceShareAssociationsReqBodyAssociationType struct {
	value string
}

type SearchResourceShareAssociationsReqBodyAssociationTypeEnum struct {
	PRINCIPAL SearchResourceShareAssociationsReqBodyAssociationType
	RESOURCE  SearchResourceShareAssociationsReqBodyAssociationType
}

func GetSearchResourceShareAssociationsReqBodyAssociationTypeEnum() SearchResourceShareAssociationsReqBodyAssociationTypeEnum {
	return SearchResourceShareAssociationsReqBodyAssociationTypeEnum{
		PRINCIPAL: SearchResourceShareAssociationsReqBodyAssociationType{
			value: "principal",
		},
		RESOURCE: SearchResourceShareAssociationsReqBodyAssociationType{
			value: "resource",
		},
	}
}

func (c SearchResourceShareAssociationsReqBodyAssociationType) Value() string {
	return c.value
}

func (c SearchResourceShareAssociationsReqBodyAssociationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchResourceShareAssociationsReqBodyAssociationType) UnmarshalJSON(b []byte) error {
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
