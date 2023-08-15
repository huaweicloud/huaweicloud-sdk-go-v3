package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterCheckBody struct {

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群规格名称
	Flavor string `json:"flavor"`

	// 可用区列表
	AvailabilityZones []string `json:"availability_zones"`

	// 实例节点个数
	NumNode int32 `json:"num_node"`

	// 集群安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 集群版本
	DatastoreVersion string `json:"datastore_version"`

	// 集群虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// 集群子网ID
	SubnetId string `json:"subnet_id"`

	PublicIp *OpenPublicIp `json:"public_ip,omitempty"`

	// 跨规格恢复
	CrossSpecRestore *string `json:"cross_spec_restore,omitempty"`

	Volume *Volume `json:"volume,omitempty"`

	// 旧主机名
	OldClusterHostname *string `json:"old_cluster_hostname,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point,omitempty"`

	// 标签列表
	TagList *[]Tag `json:"tag_list,omitempty"`

	// 存储池ID
	DssPoolId *string `json:"dss_pool_id,omitempty"`

	// 数据库端口
	DbPort *string `json:"db_port,omitempty"`

	// 管理员密码
	DbPassword *string `json:"db_password,omitempty"`

	// 管理员用户
	DbName *string `json:"db_name,omitempty"`

	// cn节点数量
	NumCn *int32 `json:"num_cn,omitempty"`

	// 集群名称
	Name *string `json:"name,omitempty"`
}

func (o ClusterCheckBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCheckBody struct{}"
	}

	return strings.Join([]string{"ClusterCheckBody", string(data)}, " ")
}
