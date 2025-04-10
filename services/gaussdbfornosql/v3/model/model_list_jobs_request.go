package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 查询开始时间，默认当前时间往前30天，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始，Z指时区偏移量
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间，默认当前时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间，时间跨度不超过30天。 其中，T指某个时间的开始，Z指时区偏移量。
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态： 取值为“Running”为执行中； 取值为“Completed”为完成； 取值为“Failed” 为失败。
	Status *string `json:"status,omitempty"`

	// 任务名称。对应取值如下： - \"CreateInstance\"：创建实例 - \"RestoreNewInstance\"：恢复到新实例 - \"EnlargeInstance\"：扩容实例 - \"ReduceInstance\"：缩容实例 - \"RestartInstance\"：重启实例 - \"RestartNode\"：重启节点 - \"EnlargeInstanceVolume\"：扩容实例磁盘 - \"ReduceInstanceVolume\"：缩容实例磁盘 - \"ResizeInstance\"：规格变更实例 - \"UpgradeDbVersion\"：升级数据库版本 - \"BindPublicIP\"：绑定公网IP - \"UnbindPublicIP\"：解绑公网IP - \"DeleteInstance\"：删除实例 - \"EnlargeInstanceColdVolume\"：扩容实例冷存储 - \"AddInstanceColdVolume\"：增加实例冷存储 - \"ModifySecurityGroup\"：修改安全组  - \"ModifyPort\"：修改端口 - \"ConstructDisasterRecovery\"：构造容灾关系 - \"DeConstructDisasterRecovery\"：解除容灾关系 - \"SwitchOverDisasterRecovery\"：切换容灾关系 - \"BuildBiActiveInstance\"：构建双活实例 - \"ReleaseBiActiveInstance\"：解除双活实例关系 - \"BackupInstance\"：备份实例
	Name *string `json:"name,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。取值 10 20 50 ，默认为50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
