package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象。
type CreateClusterBody struct {
	// 集群名称。4～32个字符，只能包含数字、字母、中划线和下划线，且必须以字母开头。

	Name string `json:"name"`

	BackupStrategy *CreateClusterBackupStrategyBody `json:"backupStrategy,omitempty"`

	Roles []CreateClusterRolesBody `json:"roles"`

	Nics *CreateClusterInstanceNicsBody `json:"nics"`
	// 企业项目ID。创建集群时，给集群绑定企业项目ID。最大长度36个字符，带\"-\"连字符的UUID格式，或者是字符串\"0\"。\"0\"表示默认企业项目。 说明：关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理服务用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 集群标签。   关于标签特性的详细信息，请参见[《标签管理产品介绍》](https://support.huaweicloud.com/productdesc-tms/zh-cn_topic_0071335169.html)。

	Tags *[]CreateClusterTagsBody `json:"tags,omitempty"`
	// 可用区。

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Datastore *CreateClusterDatastoreBody `json:"datastore"`
	// 是否开启认证，取值范围为true或false。默认关闭认证功能。当开启认证时，httpsEnable需要设置为true。  - true：表示集群开启认证。 - false：表示集群不开启认证。  此参数只有6.5.4及之后版本支持。

	AuthorityEnable *bool `json:"authorityEnable,omitempty"`
	// 设置是否进行通信加密。取值范围为true或false。默认关闭通信加密功能。当httpsEnable设置为true时，authorityEnable字段需要设置为true。  - true：表示集群进行通信加密。 - false：表示集群不进行通信加密。  此参数只有6.5.4及之后版本支持。

	HttpsEnable *bool `json:"httpsEnable,omitempty"`
	// 安全模式下集群管理员admin的密码，只有当authorityEnable设置为true时需要设置此参数。

	AdminPwd *string `json:"adminPwd,omitempty"`

	PublicIPReq *CreateClusterPublicIpReq `json:"publicIPReq,omitempty"`

	LoadBalance *CreateClusterLoadBalance `json:"loadBalance,omitempty"`

	PublicKibanaReq *CreateClusterPublicKibanaReq `json:"publicKibanaReq,omitempty"`
}

func (o CreateClusterBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterBody struct{}"
	}

	return strings.Join([]string{"CreateClusterBody", string(data)}, " ")
}
