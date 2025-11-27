package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobSpec struct {

	// Job类型
	Type *string `json:"type,omitempty"`

	// 联邦uid
	FederationUID *string `json:"federationUID,omitempty"`

	// 资源id
	ResourceID *string `json:"resourceID,omitempty"`

	// 资源名字
	ResourceName *string `json:"resourceName,omitempty"`

	// 扩展参数
	Extendparam *string `json:"extendparam,omitempty"`

	// 子Job
	SubJobs *[]Job `json:"subJobs,omitempty"`
}

func (o JobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSpec struct{}"
	}

	return strings.Join([]string{"JobSpec", string(data)}, " ")
}
