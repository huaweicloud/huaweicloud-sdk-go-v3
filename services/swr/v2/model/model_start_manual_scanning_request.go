package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartManualScanningRequest Request Object
type StartManualScanningRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 命名空间名称
	NamespaceName string `json:"namespace_name"`

	// 仓库名称
	RepositoryName string `json:"repository_name"`

	// 制品摘要
	Reference string `json:"reference"`
}

func (o StartManualScanningRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartManualScanningRequest struct{}"
	}

	return strings.Join([]string{"StartManualScanningRequest", string(data)}, " ")
}
