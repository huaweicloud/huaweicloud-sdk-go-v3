package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PutTaskReq 更新指定迁移任务
type PutTaskReq struct {

	// 任务名称（用户自定义） 仅由中文字符、下划线、短横线、数字、英文大小写字母组成
	Name *string `json:"name,omitempty"`

	// 任务类型，创建时必选，更新时可选 MIGRATE_FILE:文件级迁移 MIGRATE_BLOCK:块级迁移
	Type *PutTaskReqType `json:"type,omitempty"`

	// 目的端服务器的区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 目的端服务器的区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 目的端服务器是否存在。true代表已有目的端服务器，false代表需要新建目的端服务器
	ExistServer *bool `json:"exist_server,omitempty"`

	// 目的端服务器的IP地址。 公网迁移时请填写弹性IP地址 专线迁移时请填写私有IP地址 如果use_ipv6是true，则需要满足ipv6的格式，如果use_ipv6是false，则需要满足ipv4的格式
	MigrationIp *string `json:"migration_ip,omitempty"`

	// 目的端服务器的IP地址是否使用ipv6
	UseIpv6 *bool `json:"use_ipv6,omitempty"`

	// 是否为公网
	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	// 限制迁移速率，单位：Mbps
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 目的端服务器所在项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 目的端服务器所在项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目ID
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// 目的端服务器镜像代理磁盘ID
	ImageDiskId *string `json:"image_disk_id,omitempty"`

	// 模板ID
	VmTemplateId *string `json:"vm_template_id,omitempty"`

	// 目的端服务器磁盘ID数组，用空格分隔，单个id长度不超过255
	TargetDiskIds *string `json:"target_disk_ids,omitempty"`

	// 目的端的快照ID，id之间\";\"分隔
	SnapshotIds *string `json:"snapshot_ids,omitempty"`

	// 割接的快照ID
	CutoveredSnapshotIds *string `json:"cutovered_snapshot_ids,omitempty"`

	TargetServer *TargetServer `json:"target_server,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`
}

func (o PutTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutTaskReq struct{}"
	}

	return strings.Join([]string{"PutTaskReq", string(data)}, " ")
}

type PutTaskReqType struct {
	value string
}

type PutTaskReqTypeEnum struct {
	MIGRATE_FILE  PutTaskReqType
	MIGRATE_BLOCK PutTaskReqType
}

func GetPutTaskReqTypeEnum() PutTaskReqTypeEnum {
	return PutTaskReqTypeEnum{
		MIGRATE_FILE: PutTaskReqType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: PutTaskReqType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c PutTaskReqType) Value() string {
	return c.value
}

func (c PutTaskReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutTaskReqType) UnmarshalJSON(b []byte) error {
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
