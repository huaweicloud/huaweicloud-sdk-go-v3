package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportLocalVulInfoRequestInfo 操作的事件
type BatchExportLocalVulInfoRequestInfo struct {

	// 要导出的镜像信息列表，operate_all参数为false时需要填写批量查询条件,image_id 镜像id，唯一镜像标识（注：对私有镜像和共享镜像来说是镜像列表返回的id）
	ImageIdList *[]string `json:"image_id_list,omitempty"`

	// 若为true全量查询，可筛选条件全部查询
	OperateAll *bool `json:"operate_all,omitempty"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -app_vul : 应用漏洞
	VulType *string `json:"vul_type,omitempty"`

	// 漏洞处置状态，包含如下:   - unhandled ：未处理   - handled : 已处理
	HandleStatus *string `json:"handle_status,omitempty"`

	// **参数解释**: 扫描状态 **约束限制**: 不涉及 **取值范围**:   - unscan : 未扫描。   - success : 扫描完成。   - scanning : 扫描中。   - failed : 扫描失败。   - download_failed : 下载失败。   - image_oversized : 镜像超大。   - waiting_for_scan : 等待扫描。   - image_scan_stop : 扫描终止。 **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本
	ImageVersion *string `json:"image_version,omitempty"`

	// 镜像大小
	ImageSize *int64 `json:"image_size,omitempty"`

	// 创建时间开始日期，时间单位 毫秒（ms）
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// 创建时间结束日期，时间单位 毫秒（ms）
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// 最近一次扫描完成时间开始日期，时间单位 毫秒（ms）
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// 最近一次扫描完成时间结束日期，时间单位 毫秒（ms）
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// **参数解释**: 是否存在容器 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	HasContainer *bool `json:"has_container,omitempty"`

	// 导出镜像漏洞时的漏洞id列表
	VulIdList *[]string `json:"vul_id_list,omitempty"`
}

func (o BatchExportLocalVulInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportLocalVulInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchExportLocalVulInfoRequestInfo", string(data)}, " ")
}
