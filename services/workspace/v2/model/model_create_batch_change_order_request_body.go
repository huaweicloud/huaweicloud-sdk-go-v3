package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateBatchChangeOrderRequestBody struct {

	// 下单类型。  - ADD_VOLUME：增加磁盘  - EXTEND_VOLUME：扩容磁盘  - RESIZE：变更规格  - CHANGE_IMAGE：切换镜像  - ADD_SUB_RESOURCES：购买桌面协同资源  - DELETE_SUB_RESOURCES：退订桌面协同资源
	Type *CreateBatchChangeOrderRequestBodyType `json:"type,omitempty"`

	AddVolumeParam *EstimateAddVolumeRequestBody `json:"add_volume_param,omitempty"`

	ExtendVolumeParam *EstimateExtendVolumeRequestBody `json:"extend_volume_param,omitempty"`

	ResizeParam *CreateResizeOrderRequestBody `json:"resize_param,omitempty"`

	ChangeImageParam *CreateChangeImageOrderRequestBody `json:"change_image_param,omitempty"`

	AddSubResourcesParam *EstimateAddSubResourcesRequestBody `json:"add_sub_resources_param,omitempty"`

	DeleteSubResourcesParam *CreateDeleteSubResourcesOrderRequestBody `json:"delete_sub_resources_param,omitempty"`
}

func (o CreateBatchChangeOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchChangeOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBatchChangeOrderRequestBody", string(data)}, " ")
}

type CreateBatchChangeOrderRequestBodyType struct {
	value string
}

type CreateBatchChangeOrderRequestBodyTypeEnum struct {
	ADD_VOLUME           CreateBatchChangeOrderRequestBodyType
	EXTEND_VOLUME        CreateBatchChangeOrderRequestBodyType
	RESIZE               CreateBatchChangeOrderRequestBodyType
	CHANGE_IMAGE         CreateBatchChangeOrderRequestBodyType
	ADD_SUB_RESOURCES    CreateBatchChangeOrderRequestBodyType
	DELETE_SUB_RESOURCES CreateBatchChangeOrderRequestBodyType
}

func GetCreateBatchChangeOrderRequestBodyTypeEnum() CreateBatchChangeOrderRequestBodyTypeEnum {
	return CreateBatchChangeOrderRequestBodyTypeEnum{
		ADD_VOLUME: CreateBatchChangeOrderRequestBodyType{
			value: "ADD_VOLUME",
		},
		EXTEND_VOLUME: CreateBatchChangeOrderRequestBodyType{
			value: "EXTEND_VOLUME",
		},
		RESIZE: CreateBatchChangeOrderRequestBodyType{
			value: "RESIZE",
		},
		CHANGE_IMAGE: CreateBatchChangeOrderRequestBodyType{
			value: "CHANGE_IMAGE",
		},
		ADD_SUB_RESOURCES: CreateBatchChangeOrderRequestBodyType{
			value: "ADD_SUB_RESOURCES",
		},
		DELETE_SUB_RESOURCES: CreateBatchChangeOrderRequestBodyType{
			value: "DELETE_SUB_RESOURCES",
		},
	}
}

func (c CreateBatchChangeOrderRequestBodyType) Value() string {
	return c.value
}

func (c CreateBatchChangeOrderRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateBatchChangeOrderRequestBodyType) UnmarshalJSON(b []byte) error {
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
