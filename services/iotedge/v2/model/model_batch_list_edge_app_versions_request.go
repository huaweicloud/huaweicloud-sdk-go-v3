package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type BatchListEdgeAppVersionsRequest struct {
	// 应用版本,应用内版本唯一。

	EdgeAppId string `json:"edge_app_id"`
	// 应用版本搜索关键字

	Version *string `json:"version,omitempty"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，默认值为10，取值区间为1-1000

	Limit *int32 `json:"limit,omitempty"`
	// ai加速卡类型

	AiCardType *BatchListEdgeAppVersionsRequestAiCardType `json:"ai_card_type,omitempty"`
	// 支持架构

	Arch *BatchListEdgeAppVersionsRequestArch `json:"arch,omitempty"`
	// 应用版本状态

	State *BatchListEdgeAppVersionsRequestState `json:"state,omitempty"`
}

func (o BatchListEdgeAppVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListEdgeAppVersionsRequest struct{}"
	}

	return strings.Join([]string{"BatchListEdgeAppVersionsRequest", string(data)}, " ")
}

type BatchListEdgeAppVersionsRequestAiCardType struct {
	value string
}

type BatchListEdgeAppVersionsRequestAiCardTypeEnum struct {
	GPU         BatchListEdgeAppVersionsRequestAiCardType
	NPU         BatchListEdgeAppVersionsRequestAiCardType
	UN_EQUIPPED BatchListEdgeAppVersionsRequestAiCardType
}

func GetBatchListEdgeAppVersionsRequestAiCardTypeEnum() BatchListEdgeAppVersionsRequestAiCardTypeEnum {
	return BatchListEdgeAppVersionsRequestAiCardTypeEnum{
		GPU: BatchListEdgeAppVersionsRequestAiCardType{
			value: "GPU",
		},
		NPU: BatchListEdgeAppVersionsRequestAiCardType{
			value: "NPU",
		},
		UN_EQUIPPED: BatchListEdgeAppVersionsRequestAiCardType{
			value: "unEquipped",
		},
	}
}

func (c BatchListEdgeAppVersionsRequestAiCardType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListEdgeAppVersionsRequestAiCardType) UnmarshalJSON(b []byte) error {
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

type BatchListEdgeAppVersionsRequestArch struct {
	value string
}

type BatchListEdgeAppVersionsRequestArchEnum struct {
	X86_64 BatchListEdgeAppVersionsRequestArch
	ARM32  BatchListEdgeAppVersionsRequestArch
	ARM64  BatchListEdgeAppVersionsRequestArch
}

func GetBatchListEdgeAppVersionsRequestArchEnum() BatchListEdgeAppVersionsRequestArchEnum {
	return BatchListEdgeAppVersionsRequestArchEnum{
		X86_64: BatchListEdgeAppVersionsRequestArch{
			value: "x86_64",
		},
		ARM32: BatchListEdgeAppVersionsRequestArch{
			value: "arm32",
		},
		ARM64: BatchListEdgeAppVersionsRequestArch{
			value: "arm64",
		},
	}
}

func (c BatchListEdgeAppVersionsRequestArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListEdgeAppVersionsRequestArch) UnmarshalJSON(b []byte) error {
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

type BatchListEdgeAppVersionsRequestState struct {
	value string
}

type BatchListEdgeAppVersionsRequestStateEnum struct {
	DRAFT     BatchListEdgeAppVersionsRequestState
	PUBLISHED BatchListEdgeAppVersionsRequestState
	OFF_SHELF BatchListEdgeAppVersionsRequestState
}

func GetBatchListEdgeAppVersionsRequestStateEnum() BatchListEdgeAppVersionsRequestStateEnum {
	return BatchListEdgeAppVersionsRequestStateEnum{
		DRAFT: BatchListEdgeAppVersionsRequestState{
			value: "DRAFT",
		},
		PUBLISHED: BatchListEdgeAppVersionsRequestState{
			value: "PUBLISHED",
		},
		OFF_SHELF: BatchListEdgeAppVersionsRequestState{
			value: "OFF_SHELF",
		},
	}
}

func (c BatchListEdgeAppVersionsRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListEdgeAppVersionsRequestState) UnmarshalJSON(b []byte) error {
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
