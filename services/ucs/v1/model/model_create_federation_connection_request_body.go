package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFederationConnectionRequestBody 创建联邦网络连接请求体
type CreateFederationConnectionRequestBody struct {

	// 项目id
	ProjectID string `json:"projectID"`

	// 虚拟私有云id，必须位于上述项目中
	VpcID string `json:"vpcID"`

	// 子网的网络id，必须位于上述虚拟私有云中
	SubnetID string `json:"subnetID"`
}

func (o CreateFederationConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFederationConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFederationConnectionRequestBody", string(data)}, " ")
}
