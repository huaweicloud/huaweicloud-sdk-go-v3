package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchOperationVo 批量操作对象，只读。
type BatchOperationVo struct {

	// 批量审批ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 组ID，填写String类型替代Long类型。
	GroupId *string `json:"group_id,omitempty"`

	// 业务名。
	BizName *string `json:"biz_name,omitempty"`

	// 业务ID，填写String类型替代Long类型。
	BizId *string `json:"biz_id,omitempty"`

	// 操作结果类型枚举。RUNNING(运行中)、SUCCESS(操作成功)、FAILED(操作失败)。 枚举值：   - RUNNING: 运行中   - SUCCESS: 操作成功   - FAILED: 操作失败
	OperationStatus *BatchOperationVoOperationStatus `json:"operation_status,omitempty"`

	// 类型。
	OperationType *string `json:"operation_type,omitempty"`

	// 业务详情。
	BizInfo *string `json:"biz_info,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// remark信息。
	Remark *string `json:"remark,omitempty"`

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 操作成功个数。
	Success *int32 `json:"success,omitempty"`

	// 操作失败个数。
	Failed *int32 `json:"failed,omitempty"`

	// 当前进度。
	Rate *string `json:"rate,omitempty"`

	// 日志。
	Logs *string `json:"logs,omitempty"`

	// 分组信息。
	Groups *[]BatchOperationVo `json:"groups,omitempty"`
}

func (o BatchOperationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperationVo struct{}"
	}

	return strings.Join([]string{"BatchOperationVo", string(data)}, " ")
}

type BatchOperationVoOperationStatus struct {
	value string
}

type BatchOperationVoOperationStatusEnum struct {
	RUNNING BatchOperationVoOperationStatus
	SUCCESS BatchOperationVoOperationStatus
	FAILED  BatchOperationVoOperationStatus
}

func GetBatchOperationVoOperationStatusEnum() BatchOperationVoOperationStatusEnum {
	return BatchOperationVoOperationStatusEnum{
		RUNNING: BatchOperationVoOperationStatus{
			value: "RUNNING",
		},
		SUCCESS: BatchOperationVoOperationStatus{
			value: "SUCCESS",
		},
		FAILED: BatchOperationVoOperationStatus{
			value: "FAILED",
		},
	}
}

func (c BatchOperationVoOperationStatus) Value() string {
	return c.value
}

func (c BatchOperationVoOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchOperationVoOperationStatus) UnmarshalJSON(b []byte) error {
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
