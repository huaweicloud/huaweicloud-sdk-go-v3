package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NamespaceMetadata struct {

	// 是否公开，可选true、false
	Public NamespaceMetadataPublic `json:"public"`

	// 上传制品时，是否自动触发制品扫描，可选true、false
	AutoScan *NamespaceMetadataAutoScan `json:"auto_scan,omitempty"`

	// 是否开启制品阻断，可选true、false
	PreventVul *NamespaceMetadataPreventVul `json:"prevent_vul,omitempty"`

	// 阻断开启的场景下，如果存在漏洞，并且存在的漏洞严重程度高于此处定义的级别，则无法拉取镜像。可选值为\"none\", \"low\", \"medium\", \"high\", \"critical\"。
	Severity *NamespaceMetadataSeverity `json:"severity,omitempty"`
}

func (o NamespaceMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceMetadata struct{}"
	}

	return strings.Join([]string{"NamespaceMetadata", string(data)}, " ")
}

type NamespaceMetadataPublic struct {
	value string
}

type NamespaceMetadataPublicEnum struct {
	TRUE  NamespaceMetadataPublic
	FALSE NamespaceMetadataPublic
}

func GetNamespaceMetadataPublicEnum() NamespaceMetadataPublicEnum {
	return NamespaceMetadataPublicEnum{
		TRUE: NamespaceMetadataPublic{
			value: "true",
		},
		FALSE: NamespaceMetadataPublic{
			value: "false",
		},
	}
}

func (c NamespaceMetadataPublic) Value() string {
	return c.value
}

func (c NamespaceMetadataPublic) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NamespaceMetadataPublic) UnmarshalJSON(b []byte) error {
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

type NamespaceMetadataAutoScan struct {
	value string
}

type NamespaceMetadataAutoScanEnum struct {
	TRUE  NamespaceMetadataAutoScan
	FALSE NamespaceMetadataAutoScan
}

func GetNamespaceMetadataAutoScanEnum() NamespaceMetadataAutoScanEnum {
	return NamespaceMetadataAutoScanEnum{
		TRUE: NamespaceMetadataAutoScan{
			value: "true",
		},
		FALSE: NamespaceMetadataAutoScan{
			value: "false",
		},
	}
}

func (c NamespaceMetadataAutoScan) Value() string {
	return c.value
}

func (c NamespaceMetadataAutoScan) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NamespaceMetadataAutoScan) UnmarshalJSON(b []byte) error {
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

type NamespaceMetadataPreventVul struct {
	value string
}

type NamespaceMetadataPreventVulEnum struct {
	TRUE  NamespaceMetadataPreventVul
	FALSE NamespaceMetadataPreventVul
}

func GetNamespaceMetadataPreventVulEnum() NamespaceMetadataPreventVulEnum {
	return NamespaceMetadataPreventVulEnum{
		TRUE: NamespaceMetadataPreventVul{
			value: "true",
		},
		FALSE: NamespaceMetadataPreventVul{
			value: "false",
		},
	}
}

func (c NamespaceMetadataPreventVul) Value() string {
	return c.value
}

func (c NamespaceMetadataPreventVul) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NamespaceMetadataPreventVul) UnmarshalJSON(b []byte) error {
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

type NamespaceMetadataSeverity struct {
	value string
}

type NamespaceMetadataSeverityEnum struct {
	NONE     NamespaceMetadataSeverity
	LOW      NamespaceMetadataSeverity
	MEDIUM   NamespaceMetadataSeverity
	HIGH     NamespaceMetadataSeverity
	CRITICAL NamespaceMetadataSeverity
}

func GetNamespaceMetadataSeverityEnum() NamespaceMetadataSeverityEnum {
	return NamespaceMetadataSeverityEnum{
		NONE: NamespaceMetadataSeverity{
			value: "none",
		},
		LOW: NamespaceMetadataSeverity{
			value: "low",
		},
		MEDIUM: NamespaceMetadataSeverity{
			value: "medium",
		},
		HIGH: NamespaceMetadataSeverity{
			value: "high",
		},
		CRITICAL: NamespaceMetadataSeverity{
			value: "critical",
		},
	}
}

func (c NamespaceMetadataSeverity) Value() string {
	return c.value
}

func (c NamespaceMetadataSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NamespaceMetadataSeverity) UnmarshalJSON(b []byte) error {
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
