package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTransferAssetJobResponse Response Object
type ShowTransferAssetJobResponse struct {

	// 转移资产任务ID
	JobId *string `json:"job_id,omitempty"`

	TransferType *TransferTypeEnum `json:"transfer_type,omitempty"`

	// 转移资产列表
	TransferAssets *[]DigitalAssetSummary `json:"transfer_assets,omitempty"`

	// 任务状态 - PROCESSING: 处理过程中 - ACCEPT： 接受 - REJECT： 拒绝 - CANCEL：取消 - FAIL: 失败
	State *ShowTransferAssetJobResponseState `json:"state,omitempty"`

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

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTransferAssetJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransferAssetJobResponse struct{}"
	}

	return strings.Join([]string{"ShowTransferAssetJobResponse", string(data)}, " ")
}

type ShowTransferAssetJobResponseState struct {
	value string
}

type ShowTransferAssetJobResponseStateEnum struct {
	PROCESSING ShowTransferAssetJobResponseState
	ACCEPT     ShowTransferAssetJobResponseState
	REJECT     ShowTransferAssetJobResponseState
	CANCEL     ShowTransferAssetJobResponseState
	FAIL       ShowTransferAssetJobResponseState
}

func GetShowTransferAssetJobResponseStateEnum() ShowTransferAssetJobResponseStateEnum {
	return ShowTransferAssetJobResponseStateEnum{
		PROCESSING: ShowTransferAssetJobResponseState{
			value: "PROCESSING",
		},
		ACCEPT: ShowTransferAssetJobResponseState{
			value: "ACCEPT",
		},
		REJECT: ShowTransferAssetJobResponseState{
			value: "REJECT",
		},
		CANCEL: ShowTransferAssetJobResponseState{
			value: "CANCEL",
		},
		FAIL: ShowTransferAssetJobResponseState{
			value: "FAIL",
		},
	}
}

func (c ShowTransferAssetJobResponseState) Value() string {
	return c.value
}

func (c ShowTransferAssetJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTransferAssetJobResponseState) UnmarshalJSON(b []byte) error {
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
