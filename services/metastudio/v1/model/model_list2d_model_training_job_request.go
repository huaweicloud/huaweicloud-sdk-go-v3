package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// List2dModelTrainingJobRequest Request Object
type List2dModelTrainingJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段，支持的排序方式有： - 按创建时间排序：create_time - 按更新时间排序：update_time - 按资产排序：asset_order
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 过滤创建时间<=输入时间的记录。
	CreateUntil *string `json:"create_until,omitempty"`

	// 过滤创建时间>=输入时间的记录。
	CreateSince *string `json:"create_since,omitempty"`

	// 任务状态，默认所有状态。  可多个状态查询，使用英文逗号分隔。  如state=CREATING,PUBLISHED
	State *string `json:"state,omitempty"`

	// 查询租户id。
	QueryProjectId *string `json:"query_project_id,omitempty"`

	// 任务批次名称。
	BatchName *string `json:"batch_name,omitempty"`

	// 任务标签。
	Tag *string `json:"tag,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 分身数字人模型名称
	Name *string `json:"name,omitempty"`

	// 模型分辨率
	ModelResolution *string `json:"model_resolution,omitempty"`

	// 是否是flexus任务
	IsFlexus *bool `json:"is_flexus,omitempty"`
}

func (o List2dModelTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "List2dModelTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"List2dModelTrainingJobRequest", string(data)}, " ")
}
