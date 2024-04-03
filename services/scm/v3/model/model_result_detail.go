package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultDetail struct {

	// 证书ID。
	CertificateId string `json:"certificate_id"`

	// 当前证书在所查询服务中已部署资源总数。
	TotalNum int32 `json:"total_num"`

	// 当前证书已部署资源列表，详情请参见DeployedResourceDetail字段数据结构说明。
	DeployedResources []DeployedResourceDetail `json:"deployed_resources"`
}

func (o ResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultDetail struct{}"
	}

	return strings.Join([]string{"ResultDetail", string(data)}, " ")
}
