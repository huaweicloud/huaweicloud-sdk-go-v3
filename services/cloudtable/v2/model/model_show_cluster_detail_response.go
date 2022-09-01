package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterDetailResponse struct {

	// 集群当前状态列表： - 创建中 - 扩容中 - 重启中 - 开启opentsdb - 扩容失败 - 重启失败 - 开启opentsdb失败
	Actions *[]string `json:"actions,omitempty" xml:"actions"`

	Datastore *Datastore `json:"datastore,omitempty" xml:"datastore"`

	// 是否打开openTSDB特性。 - false：不开启 - true：开启
	EnableOpenTSDB *bool `json:"enable_openTSDB,omitempty" xml:"enable_openTSDB"`

	// 是否打开SQL查询特性。 - false：不开启 - true：开启
	EnableLemon *bool `json:"enable_lemon,omitempty" xml:"enable_lemon"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty" xml:"cluster_name"`

	// RegionServer个数。
	CuNum *string `json:"cu_num,omitempty" xml:"cu_num"`

	// TSD节点个数。
	TsdNum *string `json:"tsd_num,omitempty" xml:"tsd_num"`

	// Lemon节点个数。
	LemonNum *string `json:"lemon_num,omitempty" xml:"lemon_num"`

	// 集群底层存储类型： - OBS - HDFS
	StorageType *string `json:"storage_type,omitempty" xml:"storage_type"`

	// 集群存储配额。
	StorageQuota *string `json:"storage_quota,omitempty" xml:"storage_quota"`

	// 当前使用存储空间。
	UsedStorageSize *string `json:"used_storage_size,omitempty" xml:"used_storage_size"`

	// 是否打开IAM认证。 - false：不开启 - true：开启
	AuthMode *bool `json:"auth_mode,omitempty" xml:"auth_mode"`

	// 是否打开dfv
	EnableDfv *bool `json:"enable_dfv,omitempty" xml:"enable_dfv"`

	// 集群更新时间。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 集群创建时间。
	Created *string `json:"created,omitempty" xml:"created"`

	// 集群唯一标识，集群ID。
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id"`

	// 集群当前状态： - 200：集群正常 - 300：集群异常 - 400：集群已删除 - 303：集群创建失败
	Status *string `json:"status,omitempty" xml:"status"`

	// 内网OpenTSDB连接访问地址。
	OpenTSDBLink *string `json:"openTSDB_link,omitempty" xml:"openTSDB_link"`

	// OpenTSDB公网endpoint地址
	TsdPublicEndpoint *string `json:"tsd_public_endpoint,omitempty" xml:"tsd_public_endpoint"`

	// 内网Lemon连接访问地址。
	LemonLink *string `json:"lemon_link,omitempty" xml:"lemon_link"`

	// 内网ZooKeeper连接访问地址。
	ZookeeperLink *string `json:"zookeeper_link,omitempty" xml:"zookeeper_link"`

	// 公网HBase连接访问地址。
	HbasePublicEndpoint *string `json:"hbase_public_endpoint,omitempty" xml:"hbase_public_endpoint"`

	// 集群是否被冻结。 - false：不冻结 - true：冻结
	IsFrozen *string `json:"is_frozen,omitempty" xml:"is_frozen"`

	// VPC ID，创建集群节点所在的虚拟私有ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 子网ID，创建集群所在子网段。
	SubNetId *string `json:"sub_net_id,omitempty" xml:"sub_net_id"`

	// 安全组对应的ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 集群所属的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty" xml:"availability_zone"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailResponse", string(data)}, " ")
}
