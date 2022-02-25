package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterDetailResponse struct {
	Datastore *ClusterDetailDatastore `json:"datastore,omitempty"`
	// 节点对象列表。

	Instances *[]ClusterDetailInstances `json:"instances,omitempty"`
	// 集群上次修改时间，格式为ISO8601: CCYY-MM-DDThh:mm:ss。

	Updated *string `json:"updated,omitempty"`
	// 集群名称。

	Name *string `json:"name,omitempty"`
	// 集群创建时间，格式为ISO8601: CCYY-MM-DDThh:mm:ss。

	Created *string `json:"created,omitempty"`
	// 集群ID。

	Id *string `json:"id,omitempty"`
	// 查询返回值。  - 100：操作进行中，如创建中。 - 200：可用。 - 303：不可用，如创建失败。

	Status *string `json:"status,omitempty"`
	// 用户VPC访问IP地址和端口号。

	Endpoint *string `json:"endpoint,omitempty"`

	ActionProgress *ClusterDetailActionProgress `json:"actionProgress,omitempty"`
	// 集群当前行为集合。

	Actions *[]ActionReq `json:"actions,omitempty"`

	FailedReasons *ClusterDetailFailedReasons `json:"failed_reasons,omitempty"`
	// 是否开启认证，取值范围为true或false。默认关闭认证功能。当开启认证时，httpsEnable需要设置为true。 - true：表示集群开启认证。 - false：表示集群不开启认证。

	AuthorityEnable *bool `json:"authorityEnable,omitempty"`

	HttpsEnable *bool `json:"httpsEnable,omitempty"`
	// 集群所属的企业项目ID。  如果集群所属用户没有开通企业项目，则不会返回该参数。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 集群标签。

	Tags           *[]ClusterDetailTags `json:"tags,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailResponse", string(data)}, " ")
}
