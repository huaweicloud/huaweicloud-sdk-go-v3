package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatasourceConnectionRequestBody 创建经典型跨源连接的请求body体。
type CreateDatasourceConnectionRequestBody struct {

	// 连接名称。
	Name string `json:"name"`

	// 服务名称，目前为CloudTable.OpenTSDB/CloudTable，MRS.OPENTSDB，DWS，RDS，CSS。 说明： 大小写不敏感。
	Service string `json:"service"`

	// 用户指定安全组ID，即为需要建立连接的服务所在的安全组。
	SecurityGroupId string `json:"security_group_id"`

	// 对应服务的子网网络ID，即为需要建立连接的服务所在的子网。
	NetworkId string `json:"network_id"`

	// 对应服务对外提供的访问url。长度不能超过512个字符。
	Url string `json:"url"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateDatasourceConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatasourceConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatasourceConnectionRequestBody", string(data)}, " ")
}
