package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConfigSetRequestBody struct {

	// 配置集合的名称
	Name *string `json:"name,omitempty"`

	// 所在命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 基于Helm Chart的部署配置（当前不支持HelmRelease类型）
	HelmReleaseSpec *interface{} `json:"helmReleaseSpec,omitempty"`

	KustomizationSpec *KustomizationSpec `json:"kustomizationSpec,omitempty"`
}

func (o UpdateConfigSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigSetRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConfigSetRequestBody", string(data)}, " ")
}
