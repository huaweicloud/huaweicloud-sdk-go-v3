package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceInfo ClickHouse实例信息。
type CreateChInstanceInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 可用区。
	AzCode string `json:"az_code"`

	// 可用区模式。 取值范围： - single：单可用区 - double：多可用区
	AzMode string `json:"az_mode"`

	// 实例名。
	Name string `json:"name"`

	Engine *ClickHouseEngineInfo `json:"engine"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 数据库用户。
	DbUser string `json:"db_user"`

	// 数据库端口。取值范围：0~65535。
	Port int32 `json:"port"`

	// 部署模式。 取值范围： - Single：单机 - Ha：主备
	HaMode string `json:"ha_mode"`

	PayInfo *CreateChInstanceInfoPayInfo `json:"pay_info"`

	// SSL开关。
	SslOption bool `json:"ssl_option"`

	// 实例状态。 取值范围： - creating：创建 - normal：正常 - abnormal：异常 - createfailed：创建失败 - deleted：已删除
	Status string `json:"status"`

	// 实例所在区。
	Region string `json:"region"`

	TagsInfo *CreateChInstanceInfoTagsInfo `json:"tags_info"`
}

func (o CreateChInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceInfo struct{}"
	}

	return strings.Join([]string{"CreateChInstanceInfo", string(data)}, " ")
}
