package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 设置同步策略请求体
type SyncPolicyReq struct {
	// 任务ID。

	JobId string `json:"job_id"`
	// 冲突策略。 - ignore：忽略 - overwrite：覆盖 - stop：报错

	ConflictPolicy SyncPolicyReqConflictPolicy `json:"conflict_policy"`
	// 过滤DDL策略。

	FilterDdlPolicy SyncPolicyReqFilterDdlPolicy `json:"filter_ddl_policy"`
	// 同步增量是否同步DDL。

	DdlTrans *bool `json:"ddl_trans,omitempty"`
	// 同步增量是否同步索引。

	IndexTrans *bool `json:"index_trans,omitempty"`
}

func (o SyncPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncPolicyReq struct{}"
	}

	return strings.Join([]string{"SyncPolicyReq", string(data)}, " ")
}

type SyncPolicyReqConflictPolicy struct {
	value string
}

type SyncPolicyReqConflictPolicyEnum struct {
	IGNORE    SyncPolicyReqConflictPolicy
	OVERWRITE SyncPolicyReqConflictPolicy
	STOP      SyncPolicyReqConflictPolicy
}

func GetSyncPolicyReqConflictPolicyEnum() SyncPolicyReqConflictPolicyEnum {
	return SyncPolicyReqConflictPolicyEnum{
		IGNORE: SyncPolicyReqConflictPolicy{
			value: "ignore",
		},
		OVERWRITE: SyncPolicyReqConflictPolicy{
			value: "overwrite",
		},
		STOP: SyncPolicyReqConflictPolicy{
			value: "stop",
		},
	}
}

func (c SyncPolicyReqConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqConflictPolicy) UnmarshalJSON(b []byte) error {
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

type SyncPolicyReqFilterDdlPolicy struct {
	value string
}

type SyncPolicyReqFilterDdlPolicyEnum struct {
	DROP_DATABASE SyncPolicyReqFilterDdlPolicy
}

func GetSyncPolicyReqFilterDdlPolicyEnum() SyncPolicyReqFilterDdlPolicyEnum {
	return SyncPolicyReqFilterDdlPolicyEnum{
		DROP_DATABASE: SyncPolicyReqFilterDdlPolicy{
			value: "drop_database",
		},
	}
}

func (c SyncPolicyReqFilterDdlPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SyncPolicyReqFilterDdlPolicy) UnmarshalJSON(b []byte) error {
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
