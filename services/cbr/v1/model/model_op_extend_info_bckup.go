package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpExtendInfoBckup struct {

	// 应用一致性备份失败错误码。请参见[错误码](ErrorCode.xml)。
	AppConsistencyErrorCode *string `json:"app_consistency_error_code,omitempty" xml:"app_consistency_error_code"`

	// 应用一致性备份错误信息
	AppConsistencyErrorMessage *string `json:"app_consistency_error_message,omitempty" xml:"app_consistency_error_message"`

	// 应用一致性备份状态；0:非应用一致性，1：应用一致性备份
	AppConsistencyStatus *OpExtendInfoBckupAppConsistencyStatus `json:"app_consistency_status,omitempty" xml:"app_consistency_status"`

	// 备份副本ID
	BackupId string `json:"backup_id" xml:"backup_id"`

	// 备份名称
	BackupName *string `json:"backup_name,omitempty" xml:"backup_name"`

	// 是否增备
	Incremental *OpExtendInfoBckupIncremental `json:"incremental,omitempty" xml:"incremental"`
}

func (o OpExtendInfoBckup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoBckup struct{}"
	}

	return strings.Join([]string{"OpExtendInfoBckup", string(data)}, " ")
}

type OpExtendInfoBckupAppConsistencyStatus struct {
	value string
}

type OpExtendInfoBckupAppConsistencyStatusEnum struct {
	E_0 OpExtendInfoBckupAppConsistencyStatus
	E_1 OpExtendInfoBckupAppConsistencyStatus
}

func GetOpExtendInfoBckupAppConsistencyStatusEnum() OpExtendInfoBckupAppConsistencyStatusEnum {
	return OpExtendInfoBckupAppConsistencyStatusEnum{
		E_0: OpExtendInfoBckupAppConsistencyStatus{
			value: "0",
		},
		E_1: OpExtendInfoBckupAppConsistencyStatus{
			value: "1",
		},
	}
}

func (c OpExtendInfoBckupAppConsistencyStatus) Value() string {
	return c.value
}

func (c OpExtendInfoBckupAppConsistencyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpExtendInfoBckupAppConsistencyStatus) UnmarshalJSON(b []byte) error {
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

type OpExtendInfoBckupIncremental struct {
	value string
}

type OpExtendInfoBckupIncrementalEnum struct {
	TRUE  OpExtendInfoBckupIncremental
	FALSE OpExtendInfoBckupIncremental
}

func GetOpExtendInfoBckupIncrementalEnum() OpExtendInfoBckupIncrementalEnum {
	return OpExtendInfoBckupIncrementalEnum{
		TRUE: OpExtendInfoBckupIncremental{
			value: "\"true\"",
		},
		FALSE: OpExtendInfoBckupIncremental{
			value: "\"false\"",
		},
	}
}

func (c OpExtendInfoBckupIncremental) Value() string {
	return c.value
}

func (c OpExtendInfoBckupIncremental) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpExtendInfoBckupIncremental) UnmarshalJSON(b []byte) error {
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
