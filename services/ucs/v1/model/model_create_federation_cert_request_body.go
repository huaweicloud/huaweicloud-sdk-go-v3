package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationCertRequestBody 下载联邦kubeconfig请求体
type CreateFederationCertRequestBody struct {

	// 项目id
	ProjectID string `json:"projectID"`

	// VPC id，必须属于上述项目
	VpcID string `json:"vpcID"`

	// 子网id，必须属于上述vpc
	SubnetID string `json:"subnetID"`

	// kubeconfig证书有效期，单位为天
	Duration int32 `json:"duration"`
}

func (o CreateFederationCertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationCertRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFederationCertRequestBody", string(data)}, " ")
}
