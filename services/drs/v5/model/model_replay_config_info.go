package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReplayConfigInfo 仿真参数信息。
type ReplayConfigInfo struct {

	// 源实例ID。
	DbInstanceId string `json:"db_instance_id"`

	// 云类型： - AWSCloud：亚马逊云。 - TencentCloud：腾讯云。 - AliCloud：阿里云。
	CloudType ReplayConfigInfoCloudType `json:"cloud_type"`

	// 其他云AK信息。
	Ak string `json:"ak"`

	// 其他云SK信息。
	Sk string `json:"sk"`

	// 源数据库来源。取值： - aws_aurora_mysql：Amazon Aurora MySQL。 - tencent_tdsql_c：腾讯云TDSQL-C MySQL。 - ali_rds_mysql：阿里云RDS MySQL。
	DbSource ReplayConfigInfoDbSource `json:"db_source"`

	// 其他云Region名称。
	Region string `json:"region"`

	// 流量文件来源。 - sdk：通过三方云API接口方式下载审计日志。
	TrafficSource string `json:"traffic_source"`
}

func (o ReplayConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayConfigInfo struct{}"
	}

	return strings.Join([]string{"ReplayConfigInfo", string(data)}, " ")
}

type ReplayConfigInfoCloudType struct {
	value string
}

type ReplayConfigInfoCloudTypeEnum struct {
	AWS_CLOUD     ReplayConfigInfoCloudType
	TENCENT_CLOUD ReplayConfigInfoCloudType
	ALI_CLOUD     ReplayConfigInfoCloudType
}

func GetReplayConfigInfoCloudTypeEnum() ReplayConfigInfoCloudTypeEnum {
	return ReplayConfigInfoCloudTypeEnum{
		AWS_CLOUD: ReplayConfigInfoCloudType{
			value: "AWSCloud",
		},
		TENCENT_CLOUD: ReplayConfigInfoCloudType{
			value: "TencentCloud",
		},
		ALI_CLOUD: ReplayConfigInfoCloudType{
			value: "AliCloud",
		},
	}
}

func (c ReplayConfigInfoCloudType) Value() string {
	return c.value
}

func (c ReplayConfigInfoCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplayConfigInfoCloudType) UnmarshalJSON(b []byte) error {
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

type ReplayConfigInfoDbSource struct {
	value string
}

type ReplayConfigInfoDbSourceEnum struct {
	AWS_AURORA_MYSQL ReplayConfigInfoDbSource
	TENCENT_TDSQL_C  ReplayConfigInfoDbSource
	ALI_RDS_MYSQL    ReplayConfigInfoDbSource
}

func GetReplayConfigInfoDbSourceEnum() ReplayConfigInfoDbSourceEnum {
	return ReplayConfigInfoDbSourceEnum{
		AWS_AURORA_MYSQL: ReplayConfigInfoDbSource{
			value: "aws_aurora_mysql",
		},
		TENCENT_TDSQL_C: ReplayConfigInfoDbSource{
			value: "tencent_tdsql_c",
		},
		ALI_RDS_MYSQL: ReplayConfigInfoDbSource{
			value: "ali_rds_mysql",
		},
	}
}

func (c ReplayConfigInfoDbSource) Value() string {
	return c.value
}

func (c ReplayConfigInfoDbSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplayConfigInfoDbSource) UnmarshalJSON(b []byte) error {
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
