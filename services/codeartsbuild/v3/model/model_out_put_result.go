package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutPutResult 产物信息
type OutPutResult struct {

	// 构建任务所在项目的ID
	ProjectId *string `json:"project_id,omitempty"`

	// 产物名称
	Name *string `json:"name,omitempty"`

	// 产物版本
	Version *string `json:"version,omitempty"`

	// 产物类型
	PackageType *string `json:"package_type,omitempty"`

	// 产物路径
	Uri *string `json:"uri,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 构建编号，每日从1开始
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`
}

func (o OutPutResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutPutResult struct{}"
	}

	return strings.Join([]string{"OutPutResult", string(data)}, " ")
}
