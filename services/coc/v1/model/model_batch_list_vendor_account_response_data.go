package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BatchListVendorAccountResponseData struct {

	// **参数解释：** CMDB分配的云厂商账户ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 云广商信息。 **取值范围：** - RMS： 华为云。 - AWS：亚马逊。 - AZURE：微软。 - ALI：阿里云。 - VMWARE：VMware。 - OPENSTACK：openstack云平台。 - HCS：Huawei Cloud Stack。 - OTHER：其他云广商。
	Vendor *string `json:"vendor,omitempty"`

	// **参数解释：** 供应商的账户ID。 **取值范围：** 字符串，长度0到64个字符。
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 账户名。 **取值范围：** 字符串，长度0到64个字符。
	AccountName *string `json:"account_name,omitempty"`

	// **参数解释：** 账户ak。 **取值范围：** 字符串，长度0到64个字符。
	Ak *string `json:"ak,omitempty"`

	// **参数解释：** 任务状态。 **取值范围：** - waiting：待启动。 - running：同步中。 - success：同步成功。 - failed：同步失败。
	SyncStatus *BatchListVendorAccountResponseDataSyncStatus `json:"sync_status,omitempty"`

	// **参数解释：** 错误信息。 **取值范围：** 不涉及。
	FailureMsg *string `json:"failure_msg,omitempty"`

	// **参数解释：** 同步时间。 **取值范围：** 不涉及。
	SyncDate *sdktime.SdkTime `json:"sync_date,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 不涉及。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o BatchListVendorAccountResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListVendorAccountResponseData struct{}"
	}

	return strings.Join([]string{"BatchListVendorAccountResponseData", string(data)}, " ")
}

type BatchListVendorAccountResponseDataSyncStatus struct {
	value string
}

type BatchListVendorAccountResponseDataSyncStatusEnum struct {
	WAITING BatchListVendorAccountResponseDataSyncStatus
	RUNNING BatchListVendorAccountResponseDataSyncStatus
	SUCCESS BatchListVendorAccountResponseDataSyncStatus
	FAILED  BatchListVendorAccountResponseDataSyncStatus
}

func GetBatchListVendorAccountResponseDataSyncStatusEnum() BatchListVendorAccountResponseDataSyncStatusEnum {
	return BatchListVendorAccountResponseDataSyncStatusEnum{
		WAITING: BatchListVendorAccountResponseDataSyncStatus{
			value: "waiting",
		},
		RUNNING: BatchListVendorAccountResponseDataSyncStatus{
			value: "running",
		},
		SUCCESS: BatchListVendorAccountResponseDataSyncStatus{
			value: "success",
		},
		FAILED: BatchListVendorAccountResponseDataSyncStatus{
			value: "failed",
		},
	}
}

func (c BatchListVendorAccountResponseDataSyncStatus) Value() string {
	return c.value
}

func (c BatchListVendorAccountResponseDataSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListVendorAccountResponseDataSyncStatus) UnmarshalJSON(b []byte) error {
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
