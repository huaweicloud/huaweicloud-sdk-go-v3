package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AssetJobReason struct {

	// 资源类型
	ResourceType *AssetJobReasonResourceType `json:"resource_type,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o AssetJobReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetJobReason struct{}"
	}

	return strings.Join([]string{"AssetJobReason", string(data)}, " ")
}

type AssetJobReasonResourceType struct {
	value string
}

type AssetJobReasonResourceTypeEnum struct {
	APPLICATION AssetJobReasonResourceType
	TASK        AssetJobReasonResourceType
}

func GetAssetJobReasonResourceTypeEnum() AssetJobReasonResourceTypeEnum {
	return AssetJobReasonResourceTypeEnum{
		APPLICATION: AssetJobReasonResourceType{
			value: "application",
		},
		TASK: AssetJobReasonResourceType{
			value: "task",
		},
	}
}

func (c AssetJobReasonResourceType) Value() string {
	return c.value
}

func (c AssetJobReasonResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetJobReasonResourceType) UnmarshalJSON(b []byte) error {
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
