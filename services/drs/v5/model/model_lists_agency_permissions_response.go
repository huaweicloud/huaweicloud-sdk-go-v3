package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListsAgencyPermissionsResponse Response Object
type ListsAgencyPermissionsResponse struct {

	// - DRS FullAccess 数据复制服务所有权限
	CommonPermissions *[]string `json:"common_permissions,omitempty"`

	// -  GaussDB ReadOnlyAccess 云数据库 GaussDB服务的只读访问权限 -  GeminiDB ReadOnlyAccess 分布式多模NoSQL数据库服务只读权限 -  GaussDBforMSQLReadOnlyAccess 云数据库HUAWEIGaussDBforMSQL服务的只读访问权限 -  DWS ReadOnlyAccess 数据仓库服务只读权限 -  DDM ReadOnlyAccess 分布式数据库中间件服务只读权限 -  DDS ReadOnlyPolicy 文档数据库服务资源只读权限 -  RDS ReadOnlyAccess 关系型数据库服务资源只读权限 -  MRS ReadOnlyAccess MapReduce服务只读权限,包括集群查询操作,基础服务弹性云服务器、裸金属服务器、云硬盘、虚拟私有云只读权限
	EnginePermissions *[]string `json:"engine_permissions,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o ListsAgencyPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListsAgencyPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListsAgencyPermissionsResponse", string(data)}, " ")
}
