package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineDtoParams 源参数
type RunPipelineDtoParams struct {

	// 代码仓类型
	GitType *string `json:"git_type,omitempty"`

	// codehub代码库ID
	CodehubId *string `json:"codehub_id,omitempty"`

	// 默认分支
	DefaultBranch *string `json:"default_branch,omitempty"`

	// git仓库https地址
	GitUrl *string `json:"git_url,omitempty"`

	// 扩展点ID
	EndpointId *string `json:"endpoint_id,omitempty"`

	BuildParams *RunPipelineDtoParamsBuildParams `json:"build_params,omitempty"`
}

func (o RunPipelineDtoParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineDtoParams struct{}"
	}

	return strings.Join([]string{"RunPipelineDtoParams", string(data)}, " ")
}
