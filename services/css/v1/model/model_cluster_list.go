package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象。
type ClusterList struct {
	Datastore *ClusterListDatastore `json:"datastore,omitempty"`
	// 节点对象列表。

	Instances *[]ClusterListInstances `json:"instances,omitempty"`
	// 集群上次修改时间，格式为ISO8601: CCYY-MM-DDThh:mm:ss。

	Updated *string `json:"updated,omitempty"`
	// 集群名称。

	Name *string `json:"name,omitempty"`
	// 集群创建时间，格式为ISO8601: CCYY-MM-DDThh:mm:ss。     说明：返回的集群列表信息按照创建时间降序排序，即创建时间最新的集群排在最前。

	Created *string `json:"created,omitempty"`
	// 集群ID。

	Id *string `json:"id,omitempty"`
	// 查询返回值。  - 100：创建中。 - 200：可用。 - 303：不可用，如创建失败。

	Status *string `json:"status,omitempty"`
	// 用户VPC访问IP地址和端口号。

	Endpoint *string `json:"endpoint,omitempty"`

	ActionProgress *ClusterListActionProgress `json:"actionProgress,omitempty"`
	// 集群当前行为，REBOOTING表示重启，GROWING表示扩容，RESTORING表示恢复集群，SNAPSHOTTING表示创建快照。

	Actions *[]string `json:"actions,omitempty"`

	FailedReasons *ClusterListFailedReasons `json:"failed_reasons,omitempty"`
	// 是否开启认证，取值范围为true或false。默认关闭认证功能。当开启认证时，httpsEnable需要设置为true。 - true：表示集群开启认证。 - false：表示集群不开启认证。

	AuthorityEnable *bool `json:"authorityEnable,omitempty"`
	// VPC ID。

	VpcId *string `json:"vpcId,omitempty"`
	// 子网ID。

	SubnetId *string `json:"subnetId,omitempty"`
	// 安全组ID。

	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	// 集群所属的企业项目ID。  如果集群所属用户没有开通企业项目，则不会返回该参数。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Period *bool `json:"period,omitempty"`
	// 是否开启https访问

	HttpsEnable *bool `json:"httpsEnable,omitempty"`
	// 集群标签。

	Tags *[]ClusterListTags `json:"tags,omitempty"`
}

func (o ClusterList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterList struct{}"
	}

	return strings.Join([]string{"ClusterList", string(data)}, " ")
}
