package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteRelationResource 删除时预检查关联项
type DeleteRelationResource struct {

	// 关联资源类型,DPE,ALERT_RULE,DATAOBJECT
	Type *DeleteRelationResourceType `json:"type,omitempty"`

	// 关联数
	Count *int32 `json:"count,omitempty"`

	// 关联资源链接
	DetailUrl *string `json:"detail_url,omitempty"`

	// 关联资源信息
	Data *[]Data `json:"data,omitempty"`
}

func (o DeleteRelationResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRelationResource struct{}"
	}

	return strings.Join([]string{"DeleteRelationResource", string(data)}, " ")
}

type DeleteRelationResourceType struct {
	value string
}

type DeleteRelationResourceTypeEnum struct {
	DPE        DeleteRelationResourceType
	ALERT_RULE DeleteRelationResourceType
	DATAOBJECT DeleteRelationResourceType
}

func GetDeleteRelationResourceTypeEnum() DeleteRelationResourceTypeEnum {
	return DeleteRelationResourceTypeEnum{
		DPE: DeleteRelationResourceType{
			value: "DPE",
		},
		ALERT_RULE: DeleteRelationResourceType{
			value: "ALERT_RULE",
		},
		DATAOBJECT: DeleteRelationResourceType{
			value: "DATAOBJECT",
		},
	}
}

func (c DeleteRelationResourceType) Value() string {
	return c.value
}

func (c DeleteRelationResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteRelationResourceType) UnmarshalJSON(b []byte) error {
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
