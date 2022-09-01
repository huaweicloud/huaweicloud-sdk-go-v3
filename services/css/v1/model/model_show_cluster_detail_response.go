package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterDetailResponse struct {
	Datastore *ClusterDetailDatastore `json:"datastore,omitempty" xml:"datastore"`

	// 节点对象列表。
	Instances *[]ClusterDetailInstances `json:"instances,omitempty" xml:"instances"`

	PublicKibanaResp *PublicKibanaRespBody `json:"publicKibanaResp,omitempty" xml:"publicKibanaResp"`

	ElbWhiteList *ElbWhiteListResp `json:"elbWhiteList,omitempty" xml:"elbWhiteList"`

	// 集群上次修改时间，格式为ISO8601： CCYY-MM-DDThh:mm:ss。
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 集群名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 公网IP信息。
	PublicIp *string `json:"publicIp,omitempty" xml:"publicIp"`

	// 集群创建时间，格式为ISO8601： CCYY-MM-DDThh:mm:ss。
	Created *string `json:"created,omitempty" xml:"created"`

	// 集群ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 集群状态值。  - 100：操作进行中，如创建中。 - 200：可用。 - 303：不可用，如创建失败。
	Status *string `json:"status,omitempty" xml:"status"`

	// 用户VPC访问IP地址和端口号。
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint"`

	// VPC ID。
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId"`

	// 子网ID。
	SubnetId *string `json:"subnetId,omitempty" xml:"subnetId"`

	// 安全组ID。
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId"`

	// 终端节点IP。
	VpcepIp *string `json:"vpcepIp,omitempty" xml:"vpcepIp"`

	// 公网带宽大小。单位：Mbit/s
	BandwidthSize *int32 `json:"bandwidthSize,omitempty" xml:"bandwidthSize"`

	// 通信加密状态。 - false：未设置通信加密。 - true：已设置通信加密。
	HttpsEnable *bool `json:"httpsEnable,omitempty" xml:"httpsEnable"`

	// 磁盘是否加密。  - true : 磁盘已加密。 - false : 磁盘未加密。
	DiskEncrypted *bool `json:"diskEncrypted,omitempty" xml:"diskEncrypted"`

	// 是否开启认证，取值范围为true或false。默认关闭认证功能。 - true：表示集群开启认证。 - false：表示集群不开启认证。
	AuthorityEnable *bool `json:"authorityEnable,omitempty" xml:"authorityEnable"`

	// 快照是否开启。 - true: 快照开启状态。 - false: 快照关闭状态。
	BackupAvailable *bool `json:"backupAvailable,omitempty" xml:"backupAvailable"`

	// 集群行为进度，显示创建或扩容进度的百分比。
	ActionProgress *interface{} `json:"actionProgress,omitempty" xml:"actionProgress"`

	// 集群当前行为。REBOOTING表示重启、GROWING表示扩容、RESTORING表示恢复集群、SNAPSHOTTING表示创建快照等。
	Actions *[]interface{} `json:"actions,omitempty" xml:"actions"`

	// 集群所属的企业项目ID。  如果集群所属用户没有开通企业项目，则不会返回该参数。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty" xml:"enterpriseProjectId"`

	// 集群标签。
	Tags *[]ClusterDetailTags `json:"tags,omitempty" xml:"tags"`

	FailedReasons *ClusterDetailFailedReasons `json:"failed_reasons,omitempty" xml:"failed_reasons"`

	// 是否为包周期集群。 - \"true\" 表示是包周期计费的集群。 - \"false\" 表示是按需计费的集群。
	Period         *bool `json:"period,omitempty" xml:"period"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailResponse", string(data)}, " ")
}
