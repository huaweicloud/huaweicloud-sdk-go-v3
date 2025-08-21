package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlResponse Response Object
type ShowControlResponse struct {

	// 控制策略ID。
	Identifier *string `json:"identifier,omitempty"`

	// 策略类别。
	Implementation *string `json:"implementation,omitempty"`

	// 实施建议。
	Guidance *string `json:"guidance,omitempty"`

	// 治理资源。
	Resource *[]string `json:"resource,omitempty"`

	// 控制策略所属服务。
	Service *string `json:"service,omitempty"`

	// 控制策略类型。包括主动性控制策略Proactive、检测性控制策略Detective、预防性控制策略Preventive。
	Behavior *string `json:"behavior,omitempty"`

	// 控制策略目标。
	ControlObjective *string `json:"control_objective,omitempty"`

	// 治理策略来自的框架。
	Framework *[]string `json:"framework,omitempty"`

	// 控制策略内容。
	Artifacts *[]Artifact `json:"artifacts,omitempty"`

	// 控制策略别名。
	Aliases *[]string `json:"aliases,omitempty"`

	// 控制策略来源。
	Owner *string `json:"owner,omitempty"`

	// 控制策略严重性。
	Severity *string `json:"severity,omitempty"`

	// 控制策略版本。
	Version *string `json:"version,omitempty"`

	// 控制策略发布时间。
	ReleaseDate    *sdktime.SdkTime `json:"release_date,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlResponse struct{}"
	}

	return strings.Join([]string{"ShowControlResponse", string(data)}, " ")
}
