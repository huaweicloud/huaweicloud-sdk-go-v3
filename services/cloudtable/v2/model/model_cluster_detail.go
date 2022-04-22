package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群详情信息对象。
type ClusterDetail struct {
	ActionProgress *ActionProgress `json:"action_progress,omitempty"`

	// 集群操作记录
	Actions *[]string `json:"actions,omitempty"`

	// 是否开启IAM权限认证。 - false：不开启 - true：开启
	AuthMode *string `json:"auth_mode,omitempty"`

	// 集群所在的可用区（AZ)。
	AzCode *string `json:"az_code,omitempty"`

	// 集群ID，集群唯一标识。
	ClusterId *string `json:"cluster_id,omitempty"`

	// CloudTable集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群创建时间。
	Created *string `json:"created,omitempty"`

	// 是否开启DFV。 - false：不开启 - true：开启
	EnableDfv *string `json:"enable_dfv,omitempty"`

	// 集群是否免费。 - false：不免费 - true：免费
	EnableFree *string `json:"enable_free,omitempty"`

	// 是否开启Lemon。 - false：不开启 - true：开启
	EnableLemon *string `json:"enable_lemon,omitempty"`

	// 是否开启OpenTSDB。 - false：不开启 - true：开启
	EnableOpenTSDB *string `json:"enable_openTSDB,omitempty"`

	// 集群状态： - 200：集群正常 - 300：集群异常 - 303：集群创建失败 - 400：集群已删除
	Status *string `json:"status,omitempty"`

	// 集群标识符。
	Tags *string `json:"tags,omitempty"`

	// 集群版本号。
	Version *string `json:"version,omitempty"`

	// CloudTable集群ZooKeeper的链接地址。例如：cloudtable-3058-zk3-Dqcwuh6R.mycloudtable.com:2181,cloudtable-3058-zk2-TCwkZEie.mycloudtable.com:2181,cloudtable-3058-zk1-TBELUFOK.mycloudtable.com:2181
	ZookeeperLink *string `json:"zookeeper_link,omitempty"`
}

func (o ClusterDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetail struct{}"
	}

	return strings.Join([]string{"ClusterDetail", string(data)}, " ")
}
