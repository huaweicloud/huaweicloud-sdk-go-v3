package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesRequest Request Object
type ListVolumesRequest struct {

	// 通过云硬盘ID进行分页查询。默认为查询第一页数据。
	Marker *string `json:"marker,omitempty"`

	// 磁盘名称。
	Name *string `json:"name,omitempty"`

	// 返回结果个数限制。默认值为1000。
	Limit *int32 `json:"limit,omitempty"`

	// 返回结果按该关键字排序，支持id，status，size，created_at等关键字，默认为“created_at”。
	SortKey *string `json:"sort_key,omitempty"`

	// 偏移量（偏移量为一个大于0小于磁盘总个数的整数，表示查询该偏移量后面的所有的磁盘）。
	Offset *int32 `json:"offset,omitempty"`

	// 返回结果按照降序或升序排列，默认为“desc”。 降序：desc 升序：asc
	SortDir *string `json:"sort_dir,omitempty"`

	// 云硬盘状态，取值可参考：\"[云硬盘状态](https://support.huaweicloud.com/api-evs/evs_04_0040.html)\"。
	Status *string `json:"status,omitempty"`

	// 云硬盘元数据。
	Metadata *string `json:"metadata,omitempty"`

	// 可用区信息。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 是否为共享云硬盘。 true：表示为共享云硬盘。 false：表示为非共享云硬盘。
	Multiattach *bool `json:"multiattach,omitempty"`

	// 服务类型，仅支持EVS、DSS、DESS。
	ServiceType *string `json:"service_type,omitempty"`

	// 专属存储池ID，可过滤出该专属存储池下的所有云硬盘，必须精确匹配。
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	// 专属存储池的名字，可过滤出该专属存储池下的所有云硬盘，支持模糊匹配。
	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	// 云硬盘类型id。 通过\"[查询云硬盘类型列表](https://support.huaweicloud.com/api-evs/evs_04_3035.html)\"可以查到，即volume_types参数说明表格中的“id”
	VolumeTypeId *string `json:"volume_type_id,omitempty"`

	// 云硬盘ID。
	Id *string `json:"id,omitempty"`

	// 云硬盘id列表，格式为ids=['id1','id2',...,'idx']，返回“ids”中有效id的云硬盘详情，无效的id会被忽略。 支持查询最多60个id对应的云硬盘详情。 如果“id”和“ids”查询参数同时存在，“id”会被忽略。
	Ids *string `json:"ids,omitempty"`

	// 指定企业项目id进行过滤。 传入“all_granted_eps”，代表查询权限范围内的所有企业项目下的云硬盘。 > 说明： >  > 关于企业项目ID的获取及企业项目特性的详细信息，请参考：\"[企业管理用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)\"。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云服务器id。
	ServerId *string `json:"server_id,omitempty"`

	// 自动快照策略ID
	SnapshotPolicyId *string `json:"snapshot_policy_id,omitempty"`
}

func (o ListVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesRequest", string(data)}, " ")
}
