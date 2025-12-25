package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VifPeerDetection 虚拟接口对等体连通性探测实例
type VifPeerDetection struct {

	// 虚拟接口对等体连通性探测实例uuid
	Id *string `json:"id,omitempty"`

	// 虚拟接口对等体实例uuid
	VifPeerId *string `json:"vif_peer_id,omitempty"`

	// 租户id
	ProjectId *string `json:"project_id,omitempty"`

	// 账号id
	DomainId *string `json:"domain_id,omitempty"`

	// 探测开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 探测结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 探测状态，取值范围： - processing: 探测处理中 - complete: 探测完成 - error: 探测异常退出
	Status *VifPeerDetectionStatus `json:"status,omitempty"`

	// 丢包率
	LossRatio *int32 `json:"loss_ratio,omitempty"`
}

func (o VifPeerDetection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VifPeerDetection struct{}"
	}

	return strings.Join([]string{"VifPeerDetection", string(data)}, " ")
}

type VifPeerDetectionStatus struct {
	value string
}

type VifPeerDetectionStatusEnum struct {
	PROCESSING VifPeerDetectionStatus
	COMPLETE   VifPeerDetectionStatus
	ERROR      VifPeerDetectionStatus
}

func GetVifPeerDetectionStatusEnum() VifPeerDetectionStatusEnum {
	return VifPeerDetectionStatusEnum{
		PROCESSING: VifPeerDetectionStatus{
			value: "processing",
		},
		COMPLETE: VifPeerDetectionStatus{
			value: "complete",
		},
		ERROR: VifPeerDetectionStatus{
			value: "error",
		},
	}
}

func (c VifPeerDetectionStatus) Value() string {
	return c.value
}

func (c VifPeerDetectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VifPeerDetectionStatus) UnmarshalJSON(b []byte) error {
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
