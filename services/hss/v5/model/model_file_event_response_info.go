package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FileEventResponseInfo 文件具体变更信息
type FileEventResponseInfo struct {

	// id
	Id *int32 `json:"id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// status
	Status *FileEventResponseInfoStatus `json:"status,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 变更类型
	ChangeType *FileEventResponseInfoChangeType `json:"change_type,omitempty"`

	// 变更类别
	ChangeCategory *FileEventResponseInfoChangeCategory `json:"change_category,omitempty"`

	// modified hash
	AfterChange *string `json:"after_change,omitempty"`

	// hash
	BeforeChange *string `json:"before_change,omitempty"`

	// 最后变更时间
	LatestTime *int64 `json:"latest_time,omitempty"`
}

func (o FileEventResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileEventResponseInfo struct{}"
	}

	return strings.Join([]string{"FileEventResponseInfo", string(data)}, " ")
}

type FileEventResponseInfoStatus struct {
	value string
}

type FileEventResponseInfoStatusEnum struct {
	ALL     FileEventResponseInfoStatus
	TRUST   FileEventResponseInfoStatus
	UNTRUST FileEventResponseInfoStatus
	UNKNOWN FileEventResponseInfoStatus
}

func GetFileEventResponseInfoStatusEnum() FileEventResponseInfoStatusEnum {
	return FileEventResponseInfoStatusEnum{
		ALL: FileEventResponseInfoStatus{
			value: "all",
		},
		TRUST: FileEventResponseInfoStatus{
			value: "trust",
		},
		UNTRUST: FileEventResponseInfoStatus{
			value: "untrust",
		},
		UNKNOWN: FileEventResponseInfoStatus{
			value: "unknown",
		},
	}
}

func (c FileEventResponseInfoStatus) Value() string {
	return c.value
}

func (c FileEventResponseInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileEventResponseInfoStatus) UnmarshalJSON(b []byte) error {
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

type FileEventResponseInfoChangeType struct {
	value string
}

type FileEventResponseInfoChangeTypeEnum struct {
	ALL      FileEventResponseInfoChangeType
	REGISTRY FileEventResponseInfoChangeType
	FILE     FileEventResponseInfoChangeType
}

func GetFileEventResponseInfoChangeTypeEnum() FileEventResponseInfoChangeTypeEnum {
	return FileEventResponseInfoChangeTypeEnum{
		ALL: FileEventResponseInfoChangeType{
			value: "all",
		},
		REGISTRY: FileEventResponseInfoChangeType{
			value: "registry",
		},
		FILE: FileEventResponseInfoChangeType{
			value: "file",
		},
	}
}

func (c FileEventResponseInfoChangeType) Value() string {
	return c.value
}

func (c FileEventResponseInfoChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileEventResponseInfoChangeType) UnmarshalJSON(b []byte) error {
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

type FileEventResponseInfoChangeCategory struct {
	value string
}

type FileEventResponseInfoChangeCategoryEnum struct {
	ALL    FileEventResponseInfoChangeCategory
	MODIFY FileEventResponseInfoChangeCategory
	ADD    FileEventResponseInfoChangeCategory
	DELETE FileEventResponseInfoChangeCategory
}

func GetFileEventResponseInfoChangeCategoryEnum() FileEventResponseInfoChangeCategoryEnum {
	return FileEventResponseInfoChangeCategoryEnum{
		ALL: FileEventResponseInfoChangeCategory{
			value: "all",
		},
		MODIFY: FileEventResponseInfoChangeCategory{
			value: "modify",
		},
		ADD: FileEventResponseInfoChangeCategory{
			value: "add",
		},
		DELETE: FileEventResponseInfoChangeCategory{
			value: "delete",
		},
	}
}

func (c FileEventResponseInfoChangeCategory) Value() string {
	return c.value
}

func (c FileEventResponseInfoChangeCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileEventResponseInfoChangeCategory) UnmarshalJSON(b []byte) error {
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
