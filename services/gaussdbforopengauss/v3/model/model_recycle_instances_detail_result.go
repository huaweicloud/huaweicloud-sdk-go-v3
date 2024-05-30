package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecycleInstancesDetailResult struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`

	// 部署形态(Ha:主备版;Independent:独立部署;Combined:混合部署)。
	HaMode RecycleInstancesDetailResultHaMode `json:"ha_mode"`

	// 引擎名称
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本号。
	EngineVersion string `json:"engine_version"`

	// 计费模式（0：按需计费；1：包年/包月）。
	PayModel RecycleInstancesDetailResultPayModel `json:"pay_model"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt string `json:"created_at"`

	// 删除时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	DeletedAt string `json:"deleted_at"`

	// 磁盘类型。（SAS：high；SSD：ultrahigh；ESSD：essd）。
	VolumeType RecycleInstancesDetailResultVolumeType `json:"volume_type"`

	// 磁盘大小
	VolumeSize *string `json:"volume_size,omitempty"`

	// 数据vip。
	DataVip string `json:"data_vip"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 备份级别
	BackupLevel *string `json:"backup_level,omitempty"`

	// 备份ID。（指删除实例时产生备份信息中的备份ID）。
	RecycleBackupId string `json:"recycle_backup_id"`

	// 回收站备份状态。（Running：运行中；Active：有效的）。
	RecycleStatus RecycleInstancesDetailResultRecycleStatus `json:"recycle_status"`

	// 产品类型（basic：基础版；standard：标准版；enterprise：企业版）。
	Mode RecycleInstancesDetailResultMode `json:"mode"`
}

func (o RecycleInstancesDetailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstancesDetailResult struct{}"
	}

	return strings.Join([]string{"RecycleInstancesDetailResult", string(data)}, " ")
}

type RecycleInstancesDetailResultHaMode struct {
	value string
}

type RecycleInstancesDetailResultHaModeEnum struct {
	HA          RecycleInstancesDetailResultHaMode
	INDEPENDENT RecycleInstancesDetailResultHaMode
}

func GetRecycleInstancesDetailResultHaModeEnum() RecycleInstancesDetailResultHaModeEnum {
	return RecycleInstancesDetailResultHaModeEnum{
		HA: RecycleInstancesDetailResultHaMode{
			value: "Ha",
		},
		INDEPENDENT: RecycleInstancesDetailResultHaMode{
			value: "Independent",
		},
	}
}

func (c RecycleInstancesDetailResultHaMode) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultHaMode) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultPayModel struct {
	value string
}

type RecycleInstancesDetailResultPayModelEnum struct {
	E_0 RecycleInstancesDetailResultPayModel
	E_1 RecycleInstancesDetailResultPayModel
}

func GetRecycleInstancesDetailResultPayModelEnum() RecycleInstancesDetailResultPayModelEnum {
	return RecycleInstancesDetailResultPayModelEnum{
		E_0: RecycleInstancesDetailResultPayModel{
			value: "0",
		},
		E_1: RecycleInstancesDetailResultPayModel{
			value: "1",
		},
	}
}

func (c RecycleInstancesDetailResultPayModel) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultPayModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultPayModel) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultVolumeType struct {
	value string
}

type RecycleInstancesDetailResultVolumeTypeEnum struct {
	HIGH      RecycleInstancesDetailResultVolumeType
	ULTRAHIGH RecycleInstancesDetailResultVolumeType
	ESSD      RecycleInstancesDetailResultVolumeType
}

func GetRecycleInstancesDetailResultVolumeTypeEnum() RecycleInstancesDetailResultVolumeTypeEnum {
	return RecycleInstancesDetailResultVolumeTypeEnum{
		HIGH: RecycleInstancesDetailResultVolumeType{
			value: "high",
		},
		ULTRAHIGH: RecycleInstancesDetailResultVolumeType{
			value: "ultrahigh",
		},
		ESSD: RecycleInstancesDetailResultVolumeType{
			value: "essd",
		},
	}
}

func (c RecycleInstancesDetailResultVolumeType) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultVolumeType) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultRecycleStatus struct {
	value string
}

type RecycleInstancesDetailResultRecycleStatusEnum struct {
	RUNNING RecycleInstancesDetailResultRecycleStatus
	ACTIVE  RecycleInstancesDetailResultRecycleStatus
}

func GetRecycleInstancesDetailResultRecycleStatusEnum() RecycleInstancesDetailResultRecycleStatusEnum {
	return RecycleInstancesDetailResultRecycleStatusEnum{
		RUNNING: RecycleInstancesDetailResultRecycleStatus{
			value: "Running",
		},
		ACTIVE: RecycleInstancesDetailResultRecycleStatus{
			value: "Active",
		},
	}
}

func (c RecycleInstancesDetailResultRecycleStatus) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultRecycleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultRecycleStatus) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultMode struct {
	value string
}

type RecycleInstancesDetailResultModeEnum struct {
	BASIC      RecycleInstancesDetailResultMode
	STANDARD   RecycleInstancesDetailResultMode
	ENTERPRISE RecycleInstancesDetailResultMode
}

func GetRecycleInstancesDetailResultModeEnum() RecycleInstancesDetailResultModeEnum {
	return RecycleInstancesDetailResultModeEnum{
		BASIC: RecycleInstancesDetailResultMode{
			value: "basic",
		},
		STANDARD: RecycleInstancesDetailResultMode{
			value: "standard",
		},
		ENTERPRISE: RecycleInstancesDetailResultMode{
			value: "enterprise",
		},
	}
}

func (c RecycleInstancesDetailResultMode) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultMode) UnmarshalJSON(b []byte) error {
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
