package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployFactoryPackagesResponse Response Object
type DeployFactoryPackagesResponse struct {

	// 发布包信息
	DeployPackageDetails *[]DeployPackagesResponseDeployPackageDetails `json:"deploy_package_details,omitempty"`
	HttpStatusCode       int                                           `json:"-"`
}

func (o DeployFactoryPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployFactoryPackagesResponse struct{}"
	}

	return strings.Join([]string{"DeployFactoryPackagesResponse", string(data)}, " ")
}
