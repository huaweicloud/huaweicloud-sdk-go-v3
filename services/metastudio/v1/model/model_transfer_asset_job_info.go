package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TransferAssetJobInfo struct {

	// 转移资产任务ID
	JobId *string `json:"job_id,omitempty"`

	TransferType *TransferTypeEnum `json:"transfer_type,omitempty"`

	// 转移资产列表
	TransferAssets *[]DigitalAssetSummary `json:"transfer_assets,omitempty"`

	// 任务状态 - PROCESSING: 处理过程中 - ACCEPT： 接受 - REJECT： 拒绝 - CANCEL：取消 - FAIL: 失败
	State *TransferAssetJobInfoState `json:"state,omitempty"`

	// 源用户ID
	SrcProjectId *string `json:"src_project_id,omitempty"`

	// 目标用户ID
	DestProjectId *string `json:"dest_project_id,omitempty"`

	// 备注信息
	Memo *string `json:"memo,omitempty"`

	// 冻结/解冻原因。
	Reason *string `json:"reason,omitempty"`

	// 资产转移开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 资产转移完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 资产转移过期时间
	ExpireTime *string `json:"expire_time,omitempty"`

	// 资产转移时，是否需要从接收方扣减资源（产生计费话单）
	IsNeedBilling *bool `json:"is_need_billing,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o TransferAssetJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferAssetJobInfo struct{}"
	}

	return strings.Join([]string{"TransferAssetJobInfo", string(data)}, " ")
}

type TransferAssetJobInfoState struct {
	value string
}

type TransferAssetJobInfoStateEnum struct {
	PROCESSING TransferAssetJobInfoState
	ACCEPT     TransferAssetJobInfoState
	REJECT     TransferAssetJobInfoState
	CANCEL     TransferAssetJobInfoState
	FAIL       TransferAssetJobInfoState
}

func GetTransferAssetJobInfoStateEnum() TransferAssetJobInfoStateEnum {
	return TransferAssetJobInfoStateEnum{
		PROCESSING: TransferAssetJobInfoState{
			value: "PROCESSING",
		},
		ACCEPT: TransferAssetJobInfoState{
			value: "ACCEPT",
		},
		REJECT: TransferAssetJobInfoState{
			value: "REJECT",
		},
		CANCEL: TransferAssetJobInfoState{
			value: "CANCEL",
		},
		FAIL: TransferAssetJobInfoState{
			value: "FAIL",
		},
	}
}

func (c TransferAssetJobInfoState) Value() string {
	return c.value
}

func (c TransferAssetJobInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransferAssetJobInfoState) UnmarshalJSON(b []byte) error {
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
