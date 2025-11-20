package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupResponse Response Object
type ShowBackupResponse struct {

	// 备份id。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 关联DN。
	RelatedDataNodes *[]RelatedDn `json:"related_data_nodes,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupResponse", string(data)}, " ")
}
