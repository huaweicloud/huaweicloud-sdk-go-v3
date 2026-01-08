package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListItaSubJobsRequest Request Object
type ListItaSubJobsRequest struct {

	// 任务状态 - SUCCESS：成功。 - RUNNING：运行中。 - FAILED：失败。 - WAITING：等待。
	Status *string `json:"status,omitempty"`

	// 任务ID。 “job_type”与“job_id”建议至少填写一个。
	JobId *string `json:"job_id,omitempty"`

	// 任务类型。 “job_type”与“job_id”建议至少填写一个  - createDesktops：创建桌面任务。  - applyWorkspace：开通云桌面服务。  - cancelWorkspace：注销云桌面服务。  - expandVolumes: 扩容磁盘。  - addVolumes: 添加磁盘。  - rebuildDesktops：重建桌面系统盘。  - createSnapshot：创建磁盘快照。  - deleteSnapshot：删除磁盘快照。  - deleteDesktops：删除桌面。  - desktopRejoinDomain：桌面重新加域。  - operateDesktops：操作桌面。  - restoreDesktopBySnapshot：使用快照恢复桌面。  - desktopToImage：桌面转镜像。  - attachDesktops：分配桌面。  - deleteVolumes：删除桌面磁盘（数据盘）。  - createWksSnapshot：创建快照。  - deleteWksSnapshot：删除快照。    - createDesktopPool：创建桌面池。  - expandDesktopPool：扩容桌面池。   - deleteDesktopPoolVolumes：桌面池删除桌面磁盘（数据盘）。  - rebuildDesktopPool：桌面池重建系统盘。  - addDesktopPoolVolumes：桌面池添加磁盘。  - detachDesktops：解绑桌面池桌面。  - attachDesktopPool：分配桌面池桌面。  - batchDetachDesktops：批量解绑桌面用户。  - executeDesktopsAction：桌面操作。  - changeDesktopNetwork：桌面切换网络。  - batchChangeDesktopNetwork：桌面切换网络。
	JobType *string `json:"job_type,omitempty"`

	// 桌面池ID。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 用于分页查询，取值范围0~1000，默认1000。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListItaSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListItaSubJobsRequest struct{}"
	}

	return strings.Join([]string{"ListItaSubJobsRequest", string(data)}, " ")
}
