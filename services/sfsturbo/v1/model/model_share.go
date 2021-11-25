package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建文件系统参数body
type Share struct {
	// 文件系统所在可用区(az)的编码

	AvailabilityZone string `json:"availability_zone"`
	// 文件系统描述信息，长度为0~255。当前不支持。

	Description *string `json:"description,omitempty"`
	// 创建文件系统时，给文件系统绑定的企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`
	// SFS Turbo文件系统的名称。长度为4~64位，必须以字母开头，可以包含字母、数字、中划线、下划线，不能包含其他的特殊字符，不区分大小写。

	Name string `json:"name"`
	// 用户在某一区域下的安全组ID。

	SecurityGroupId string `json:"security_group_id"`
	// 文件系统共享协议，有效值为NFS。NFS（Network File System），即网络文件系统。一种使用于分散式文件系统的协议，通过网络让不同的机器、不同的操作系统能够彼此分享数据。

	ShareProto string `json:"share_proto"`
	// 文件系统类型，有效值为STANDARD或者PERFORMANCE。

	ShareType string `json:"share_type"`
	// 普通文件系统容量，单位GB，取值范围500~32768。 增强型文件系统，即在“metadata”字段中设置了expand_type字段，则容量范围是10240~327680

	Size int32 `json:"size"`
	// 用户在VPC下面的子网的网络ID。

	SubnetId string `json:"subnet_id"`
	// 用户在某一区域下的VPC ID。

	VpcId string `json:"vpc_id"`
	// 备份ID，从备份创建文件系统时为必选。

	BackupId *string `json:"backup_id,omitempty"`
}

func (o Share) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Share struct{}"
	}

	return strings.Join([]string{"Share", string(data)}, " ")
}
