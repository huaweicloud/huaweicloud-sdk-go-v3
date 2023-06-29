package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOnlineMigrationTaskBody 创建在线数据迁移任务结构体
type CreateOnlineMigrationTaskBody struct {

	// 在线迁移任务名称。
	Name string `json:"name"`

	// 在线迁移任务描述。
	Description *string `json:"description,omitempty"`

	// 虚拟私有云ID。   获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。  - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_apiv3_0003.html)。
	VpcId string `json:"vpc_id"`

	// 子网的网络ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。  - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。
	SubnetId string `json:"subnet_id"`

	// 指定实例所属的安全组。 安全组用来实现安全组内和组间虚拟机的访问控制，加强虚拟机的安全保护。您可以在安全组中定义各种访问规则，当虚拟机加入该安全组后，即受到这些访问规则的保护。   获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，访问控制安全组选项下可以对安全组进行创建和配置,并获取安全组ID。  - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询安全组列表](https://support.huaweicloud.com/api-vpc/vpc_apiv3_0012.html)。
	SecurityGroupId string `json:"security_group_id"`
}

func (o CreateOnlineMigrationTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskBody struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskBody", string(data)}, " ")
}
