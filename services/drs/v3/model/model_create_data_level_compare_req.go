package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type CreateDataLevelCompareReq struct {

	// 一个任务只允许有一个未完成的数据级对比任务，该字段决定对未完成数据级对比任务的处理方式。cancel-取消后重新创建,keep-保持未完成的不再创建。
	ConflictPolicy CreateDataLevelCompareReqConflictPolicy `json:"conflict_policy" xml:"conflict_policy"`

	// 数据级对比类型，lines-行对比,contents-内容对比。
	CompareType CreateDataLevelCompareReqCompareType `json:"compare_type" xml:"compare_type"`

	// 数据级对比模式，取值为空时需要在compare_object_infos或者compare_object_infos_with_token传对象信息，quick_comparison-快速对比，需要加入该功能的白名单才能使用。
	CompareMode *CreateDataLevelCompareReqCompareMode `json:"compare_mode,omitempty" xml:"compare_mode"`

	// 对比任务启动时间，取值为空代表立即启动。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 数据级对比的对象。非“快速对比”模式时，compare_object_infos和compare_object_infos_with_token根据链路二选一传值，不允许都为空。
	CompareObjectInfos *[]CompareObjectInfo `json:"compare_object_infos,omitempty" xml:"compare_object_infos"`

	// 数据级对比的对象（Cassandra灾备专用，带token信息）。非“快速对比”模式时，compare_object_infos和compare_object_infos_with_token根据链路二选一传值，不允许都为空。
	CompareObjectInfosWithToken *[]CompareObjectInfoWithToken `json:"compare_object_infos_with_token,omitempty" xml:"compare_object_infos_with_token"`
}

func (o CreateDataLevelCompareReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataLevelCompareReq struct{}"
	}

	return strings.Join([]string{"CreateDataLevelCompareReq", string(data)}, " ")
}

type CreateDataLevelCompareReqConflictPolicy struct {
	value string
}

type CreateDataLevelCompareReqConflictPolicyEnum struct {
	CANCEL CreateDataLevelCompareReqConflictPolicy
	KEEP   CreateDataLevelCompareReqConflictPolicy
}

func GetCreateDataLevelCompareReqConflictPolicyEnum() CreateDataLevelCompareReqConflictPolicyEnum {
	return CreateDataLevelCompareReqConflictPolicyEnum{
		CANCEL: CreateDataLevelCompareReqConflictPolicy{
			value: "cancel",
		},
		KEEP: CreateDataLevelCompareReqConflictPolicy{
			value: "keep",
		},
	}
}

func (c CreateDataLevelCompareReqConflictPolicy) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqConflictPolicy) UnmarshalJSON(b []byte) error {
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

type CreateDataLevelCompareReqCompareType struct {
	value string
}

type CreateDataLevelCompareReqCompareTypeEnum struct {
	LINES    CreateDataLevelCompareReqCompareType
	CONTENTS CreateDataLevelCompareReqCompareType
}

func GetCreateDataLevelCompareReqCompareTypeEnum() CreateDataLevelCompareReqCompareTypeEnum {
	return CreateDataLevelCompareReqCompareTypeEnum{
		LINES: CreateDataLevelCompareReqCompareType{
			value: "lines",
		},
		CONTENTS: CreateDataLevelCompareReqCompareType{
			value: "contents",
		},
	}
}

func (c CreateDataLevelCompareReqCompareType) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqCompareType) UnmarshalJSON(b []byte) error {
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

type CreateDataLevelCompareReqCompareMode struct {
	value string
}

type CreateDataLevelCompareReqCompareModeEnum struct {
	QUICK_COMPARISON CreateDataLevelCompareReqCompareMode
}

func GetCreateDataLevelCompareReqCompareModeEnum() CreateDataLevelCompareReqCompareModeEnum {
	return CreateDataLevelCompareReqCompareModeEnum{
		QUICK_COMPARISON: CreateDataLevelCompareReqCompareMode{
			value: "quick_comparison",
		},
	}
}

func (c CreateDataLevelCompareReqCompareMode) Value() string {
	return c.value
}

func (c CreateDataLevelCompareReqCompareMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelCompareReqCompareMode) UnmarshalJSON(b []byte) error {
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
