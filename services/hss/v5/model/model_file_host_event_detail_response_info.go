package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FileHostEventDetailResponseInfo 某个服务器文件具体变更信息
type FileHostEventDetailResponseInfo struct {

	// id
	Id *int32 `json:"id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件可信状态
	Status *FileHostEventDetailResponseInfoStatus `json:"status,omitempty"`

	// 变更类型
	ChangeType *FileHostEventDetailResponseInfoChangeType `json:"change_type,omitempty"`

	// 变更类别
	ChangeCategory *FileHostEventDetailResponseInfoChangeCategory `json:"change_category,omitempty"`

	// modified hash
	AfterChange *string `json:"after_change,omitempty"`

	// hash
	BeforeChange *string `json:"before_change,omitempty"`

	// 最后变更时间
	LatestTime *int64 `json:"latest_time,omitempty"`
}

func (o FileHostEventDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileHostEventDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"FileHostEventDetailResponseInfo", string(data)}, " ")
}

type FileHostEventDetailResponseInfoStatus struct {
	value string
}

type FileHostEventDetailResponseInfoStatusEnum struct {
	ALL     FileHostEventDetailResponseInfoStatus
	TRUST   FileHostEventDetailResponseInfoStatus
	UNTRUST FileHostEventDetailResponseInfoStatus
	UNKNOWN FileHostEventDetailResponseInfoStatus
}

func GetFileHostEventDetailResponseInfoStatusEnum() FileHostEventDetailResponseInfoStatusEnum {
	return FileHostEventDetailResponseInfoStatusEnum{
		ALL: FileHostEventDetailResponseInfoStatus{
			value: "all",
		},
		TRUST: FileHostEventDetailResponseInfoStatus{
			value: "trust",
		},
		UNTRUST: FileHostEventDetailResponseInfoStatus{
			value: "untrust",
		},
		UNKNOWN: FileHostEventDetailResponseInfoStatus{
			value: "unknown",
		},
	}
}

func (c FileHostEventDetailResponseInfoStatus) Value() string {
	return c.value
}

func (c FileHostEventDetailResponseInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileHostEventDetailResponseInfoStatus) UnmarshalJSON(b []byte) error {
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

type FileHostEventDetailResponseInfoChangeType struct {
	value string
}

type FileHostEventDetailResponseInfoChangeTypeEnum struct {
	ALL      FileHostEventDetailResponseInfoChangeType
	REGISTRY FileHostEventDetailResponseInfoChangeType
	FILE     FileHostEventDetailResponseInfoChangeType
}

func GetFileHostEventDetailResponseInfoChangeTypeEnum() FileHostEventDetailResponseInfoChangeTypeEnum {
	return FileHostEventDetailResponseInfoChangeTypeEnum{
		ALL: FileHostEventDetailResponseInfoChangeType{
			value: "all",
		},
		REGISTRY: FileHostEventDetailResponseInfoChangeType{
			value: "registry",
		},
		FILE: FileHostEventDetailResponseInfoChangeType{
			value: "file",
		},
	}
}

func (c FileHostEventDetailResponseInfoChangeType) Value() string {
	return c.value
}

func (c FileHostEventDetailResponseInfoChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileHostEventDetailResponseInfoChangeType) UnmarshalJSON(b []byte) error {
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

type FileHostEventDetailResponseInfoChangeCategory struct {
	value string
}

type FileHostEventDetailResponseInfoChangeCategoryEnum struct {
	ALL    FileHostEventDetailResponseInfoChangeCategory
	MODIFY FileHostEventDetailResponseInfoChangeCategory
	ADD    FileHostEventDetailResponseInfoChangeCategory
	DELETE FileHostEventDetailResponseInfoChangeCategory
}

func GetFileHostEventDetailResponseInfoChangeCategoryEnum() FileHostEventDetailResponseInfoChangeCategoryEnum {
	return FileHostEventDetailResponseInfoChangeCategoryEnum{
		ALL: FileHostEventDetailResponseInfoChangeCategory{
			value: "all",
		},
		MODIFY: FileHostEventDetailResponseInfoChangeCategory{
			value: "modify",
		},
		ADD: FileHostEventDetailResponseInfoChangeCategory{
			value: "add",
		},
		DELETE: FileHostEventDetailResponseInfoChangeCategory{
			value: "delete",
		},
	}
}

func (c FileHostEventDetailResponseInfoChangeCategory) Value() string {
	return c.value
}

func (c FileHostEventDetailResponseInfoChangeCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileHostEventDetailResponseInfoChangeCategory) UnmarshalJSON(b []byte) error {
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
