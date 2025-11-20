package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoadSchemaMetadataReq struct {

	// 逻辑库信息。
	CompressedDatabasesInfo *string `json:"compressed_databases_info,omitempty"`

	// 关联的后端DN信息。
	DnInstance *[]DnInstance `json:"dn_instance,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	IamAccount *IamAccount `json:"iam_account,omitempty"`
}

func (o LoadSchemaMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadSchemaMetadataReq struct{}"
	}

	return strings.Join([]string{"LoadSchemaMetadataReq", string(data)}, " ")
}
