package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecycleInstancesDetailResultV1 struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`

	// 部署形态(Ha:主备版;Independent:独立部署;Combined:混合部署)。
	HaMode RecycleInstancesDetailResultV1HaMode `json:"ha_mode"`

	// 引擎名称
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本号。
	EngineVersion string `json:"engine_version"`

	// 计费模式（0：按需计费；1：包年/包月）。
	PayModel RecycleInstancesDetailResultV1PayModel `json:"pay_model"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt string `json:"created_at"`

	// 删除时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	DeletedAt string `json:"deleted_at"`

	// 磁盘类型。（SAS：high；SSD：ultrahigh；ESSD：essd）。
	VolumeType RecycleInstancesDetailResultV1VolumeType `json:"volume_type"`

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
	RecycleStatus RecycleInstancesDetailResultV1RecycleStatus `json:"recycle_status"`

	// 产品类型（basic：基础版；standard：标准版；enterprise：企业版）。
	Mode RecycleInstancesDetailResultV1Mode `json:"mode"`
}

func (o RecycleInstancesDetailResultV1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstancesDetailResultV1 struct{}"
	}

	return strings.Join([]string{"RecycleInstancesDetailResultV1", string(data)}, " ")
}

type RecycleInstancesDetailResultV1HaMode struct {
	value string
}

type RecycleInstancesDetailResultV1HaModeEnum struct {
	HA          RecycleInstancesDetailResultV1HaMode
	INDEPENDENT RecycleInstancesDetailResultV1HaMode
}

func GetRecycleInstancesDetailResultV1HaModeEnum() RecycleInstancesDetailResultV1HaModeEnum {
	return RecycleInstancesDetailResultV1HaModeEnum{
		HA: RecycleInstancesDetailResultV1HaMode{
			value: "Ha",
		},
		INDEPENDENT: RecycleInstancesDetailResultV1HaMode{
			value: "Independent",
		},
	}
}

func (c RecycleInstancesDetailResultV1HaMode) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultV1HaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultV1HaMode) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultV1PayModel struct {
	value string
}

type RecycleInstancesDetailResultV1PayModelEnum struct {
	E_0 RecycleInstancesDetailResultV1PayModel
	E_1 RecycleInstancesDetailResultV1PayModel
}

func GetRecycleInstancesDetailResultV1PayModelEnum() RecycleInstancesDetailResultV1PayModelEnum {
	return RecycleInstancesDetailResultV1PayModelEnum{
		E_0: RecycleInstancesDetailResultV1PayModel{
			value: "0",
		},
		E_1: RecycleInstancesDetailResultV1PayModel{
			value: "1",
		},
	}
}

func (c RecycleInstancesDetailResultV1PayModel) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultV1PayModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultV1PayModel) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultV1VolumeType struct {
	value string
}

type RecycleInstancesDetailResultV1VolumeTypeEnum struct {
	HIGH      RecycleInstancesDetailResultV1VolumeType
	ULTRAHIGH RecycleInstancesDetailResultV1VolumeType
	ESSD      RecycleInstancesDetailResultV1VolumeType
}

func GetRecycleInstancesDetailResultV1VolumeTypeEnum() RecycleInstancesDetailResultV1VolumeTypeEnum {
	return RecycleInstancesDetailResultV1VolumeTypeEnum{
		HIGH: RecycleInstancesDetailResultV1VolumeType{
			value: "high",
		},
		ULTRAHIGH: RecycleInstancesDetailResultV1VolumeType{
			value: "ultrahigh",
		},
		ESSD: RecycleInstancesDetailResultV1VolumeType{
			value: "essd",
		},
	}
}

func (c RecycleInstancesDetailResultV1VolumeType) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultV1VolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultV1VolumeType) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultV1RecycleStatus struct {
	value string
}

type RecycleInstancesDetailResultV1RecycleStatusEnum struct {
	RUNNING RecycleInstancesDetailResultV1RecycleStatus
	ACTIVE  RecycleInstancesDetailResultV1RecycleStatus
}

func GetRecycleInstancesDetailResultV1RecycleStatusEnum() RecycleInstancesDetailResultV1RecycleStatusEnum {
	return RecycleInstancesDetailResultV1RecycleStatusEnum{
		RUNNING: RecycleInstancesDetailResultV1RecycleStatus{
			value: "Running",
		},
		ACTIVE: RecycleInstancesDetailResultV1RecycleStatus{
			value: "Active",
		},
	}
}

func (c RecycleInstancesDetailResultV1RecycleStatus) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultV1RecycleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultV1RecycleStatus) UnmarshalJSON(b []byte) error {
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

type RecycleInstancesDetailResultV1Mode struct {
	value string
}

type RecycleInstancesDetailResultV1ModeEnum struct {
	BASIC      RecycleInstancesDetailResultV1Mode
	STANDARD   RecycleInstancesDetailResultV1Mode
	ENTERPRISE RecycleInstancesDetailResultV1Mode
}

func GetRecycleInstancesDetailResultV1ModeEnum() RecycleInstancesDetailResultV1ModeEnum {
	return RecycleInstancesDetailResultV1ModeEnum{
		BASIC: RecycleInstancesDetailResultV1Mode{
			value: "basic",
		},
		STANDARD: RecycleInstancesDetailResultV1Mode{
			value: "standard",
		},
		ENTERPRISE: RecycleInstancesDetailResultV1Mode{
			value: "enterprise",
		},
	}
}

func (c RecycleInstancesDetailResultV1Mode) Value() string {
	return c.value
}

func (c RecycleInstancesDetailResultV1Mode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleInstancesDetailResultV1Mode) UnmarshalJSON(b []byte) error {
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
