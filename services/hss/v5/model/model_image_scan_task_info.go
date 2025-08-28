package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageScanTaskInfo 镜像扫描任务列表
type ImageScanTaskInfo struct {

	// **参数解释**： 任务ID **取值范围**： 字符长度1-256位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 策略ID **取值范围**： 字符长度1-256位
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 任务名称 **取值范围**： 字符长度1-256位
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 任务开始时间 **取值范围**： 最小值0，最大值9223372036854775807
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释**： 任务结束时间 **取值范围**： 最小值0，最大值9223372036854775807
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 任务剩余时间 **取值范围**： 最小值0，最大值2147483547
	RemainMin *int64 `json:"remain_min,omitempty"`

	// **参数解释**： 任务细分类型 **取值范围**： - cycle：定时扫描。 - manual：手动扫描。 - autoSync：定时同步。 - manualSync：手动同步。
	TaskType *string `json:"task_type,omitempty"`

	// **参数解释**： 镜像类型 **取值范围**： - local：本地镜像。 - registry：仓库镜像。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**： 扫描进度状态 **取值范围**： - 100：扫描完成。 - 0-100：扫描进度。 - -1：扫描终止。 - -2：扫描超时。 - -3：扫描异常。
	TaskStatus *int32 `json:"task_status,omitempty"`

	// **参数解释**： 扫描风险类型 **取值范围**: - 0：none。 - 0x7fffffff：全部。 - 0x000f0000：漏洞。 - 0x0000f000：基线检查。 - 0x00000f00：恶意文件。 - 0x000000f0：敏感信息。 - 0x0000000f：软件合规。
	ScanScope *int32 `json:"scan_scope,omitempty"`

	// **参数解释**： 扫描限速 单位：个/h **取值范围**： 最小值0，最大值1000
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// **参数解释**： 扫描全部镜像 **取值范围**： - true：扫描全部镜像。 - false：指定镜像扫描。
	IsAll *bool `json:"is_all,omitempty"`

	// **参数解释**： 扫描失败镜像数量 **取值范围**： 最小值0，最大值2147483547
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释**： 扫描成功镜像数量 **取值范围**： 最小值0，最大值2147483547
	SuccessNum *int32 `json:"success_num,omitempty"`

	// **参数解释**： 任务关联的镜像总数 **取值范围**： 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 有漏洞风险、基线风险、恶意文件的镜像总数 **取值范围**： 最小值0，最大值2147483547
	RiskyNum *int32 `json:"risky_num,omitempty"`

	// **参数解释**： 同步任务类型 **取值范围**： 字符长度1-256位
	SyncTaskType *string `json:"sync_task_type,omitempty"`

	// **参数解释**： 失败原因 **取值范围**： 字符长度0-128位
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**： 失败镜像列表 **取值范围**： 最小值0，最大值2147483547
	FailedImages *[]ImageScanTaskInfoFailedImages `json:"failed_images,omitempty"`
}

func (o ImageScanTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageScanTaskInfo struct{}"
	}

	return strings.Join([]string{"ImageScanTaskInfo", string(data)}, " ")
}
