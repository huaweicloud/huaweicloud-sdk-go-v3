package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterReqBody 创建集群发起的请求的请求体对象。
type CreateClusterReqBody struct {

	// 集群名字
	ClusterName string `json:"cluster_name"`

	Datastore *Datastore `json:"datastore"`

	// 可用区
	AvailabilityZone string `json:"availability_zone"`

	// 集群所在的网络信息以及安全组信息。
	Nics []Nic `json:"nics"`

	ClusterInfo *CreateClusterReqBodyClusterInfo `json:"cluster_info"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 虚拟私有云ID
	VpcId string `json:"vpc_id"`

	// type为doris时，必填项，管理账号：admin
	Dbuser *string `json:"dbuser,omitempty"`

	// type为doris时，必填项，管理账号密码
	Dbuserpwd *string `json:"dbuserpwd,omitempty"`
}

func (o CreateClusterReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqBody struct{}"
	}

	return strings.Join([]string{"CreateClusterReqBody", string(data)}, " ")
}
