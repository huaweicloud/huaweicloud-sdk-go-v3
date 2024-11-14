package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Share 创建文件系统参数body
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

	// 文件系统类型，有效值为STANDARD或者PERFORMANCE。当文件系统正在创建时，该字段不返回。  - SFS Turbo上一代文件系统规格类型：标准型和标准型增强版填写STANDARD，性能型和性能型增强版填写PERFORMANCE。  - 20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/TiB：不校验该字段，填写STANDARD或者PERFORMANCE。  - HPC缓存型：不校验该字段，填写STANDARD或者PERFORMANCE。
	ShareType string `json:"share_type"`

	// - SFS Turbo上一代文件系统规格类型-文件系统容量：取值范围为500~32768，单位为GiB。  - SFS Turbo上一代文件系统规格类型-增强版文件系统：在“metadata”字段中设置了expand_type=\"bandwidth\"，则容量范围为10240~327680，单位为GiB。  - 20MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"20M\"，则容量范围为3686~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - 40MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"40M\"，则容量范围为1228~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - 125MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"125M\"，则容量范围为1228~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - 250MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"250M\"，则容量范围为1228~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - 500MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"500M\"，则容量范围为1228~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - 1000MB/s/TiB：在“metadata”字段中设置了expand_type=\"hpc\"、hpc_bw=\"1000M\"，则容量范围为1228~1048576，单位为GiB。容量必须为1.2TiB的倍数，换算为GiB后需要向下取整。如3.6TiB->3686GiB, 4.8TiB->4915GiB，8.4TiB->8601GiB。  - HPC缓存型文件系统：在“metadata”字段中设置了expand_type=\"hpc_cache\"，则容量范围为4096~1048576，单位为GiB。不同带宽，起步容量不一样，步长均为1TiB。如2GB/s带宽，起步容量为4TiB，即4096GiB；4GB/s带宽，起步容量为8TiB，即8192GiB；8GB/s带宽，起步容量为16TiB，即16384GiB。
	Size int32 `json:"size"`

	// 用户在VPC下面的子网的网络ID。
	SubnetId string `json:"subnet_id"`

	// 用户在某一区域下的VPC ID。
	VpcId string `json:"vpc_id"`

	// 备份ID，从备份创建文件系统时为必选。
	BackupId *string `json:"backup_id,omitempty"`

	// tag标签的列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o Share) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Share struct{}"
	}

	return strings.Join([]string{"Share", string(data)}, " ")
}
