package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsRequest Request Object
type ListSnapshotsRequest struct {

	// 偏移量。 说明:分页查询快照时使用，与limit配合使用。假如共有30个快照，设置offset为11，limit为10，即为从第12个快照开始查询，一次最多可读取10个快照。
	Offset *int32 `json:"offset,omitempty"`

	// 返回结果个数限制，值为大于0的整数。默认值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 云硬盘快照名称。最大支持255个字节。
	Name *string `json:"name,omitempty"`

	// 云硬盘快照状态，具体请参见A.3 云硬盘快照状态。
	Status *string `json:"status,omitempty"`

	// 快照所属云硬盘的ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 快照所属云硬盘的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 指定快照id进行过滤。可以传入多个id过滤查询，格式：id=id1&id=id2&id=id3
	Id *string `json:"id,omitempty"`

	// 专属存储的名称。
	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	// 专属存储ID。
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	// 服务类型。仅支持EVS、DSS、DESS。
	ServiceType *string `json:"service_type,omitempty"`

	// 指定企业项目id进行过滤。 传入“all_granted_eps”，代表查询权限范围内的所有企业项目下的云硬盘。 > 说明： >  > 关于企业项目ID的获取及企业项目特性的详细信息，请参考：\"[企业管理用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)\"。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotsRequest", string(data)}, " ")
}
