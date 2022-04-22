package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象。
type CreateClusterBody struct {

	// 集群名称。4～32个字符，只能包含数字、字母、中划线和下划线，且必须以字母开头。
	Name string `json:"name"`

	BackupStrategy *CreateClusterBackupStrategyBody `json:"backupStrategy,omitempty"`

	// 集群实例个数，取值范围为1~32。
	InstanceNum int32 `json:"instanceNum"`

	Instance *CreateClusterInstanceBody `json:"instance"`

	// 企业项目ID。创建集群时，给集群绑定企业项目ID。最大长度36个字符，带\"-\"连字符的UUID格式，或者是字符串\"0\"。\"0\"表示默认企业项目。  说明：关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理服务用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群标签。 关于标签特性的详细信息，请参见[《标签管理产品介绍》](https://support.huaweicloud.com/productdesc-tms/zh-cn_topic_0071335169.html)。
	Tags *[]CreateClusterTagsBody `json:"tags,omitempty"`

	Datastore *CreateClusterDatastoreBody `json:"datastore"`

	PayInfo *PayInfoBody `json:"payInfo,omitempty"`
}

func (o CreateClusterBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterBody struct{}"
	}

	return strings.Join([]string{"CreateClusterBody", string(data)}, " ")
}
