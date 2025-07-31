package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileStatisticResponse Response Object
type ShowFileStatisticResponse struct {

	// 服务器总数
	HostNum *int32 `json:"host_num,omitempty"`

	// 总变更数
	ChangeTotalNum *int32 `json:"change_total_num,omitempty"`

	// 变更文件数
	ChangeFileNum *int32 `json:"change_file_num,omitempty"`

	// 变更注册表数量
	ChangeRegistryNum *int32 `json:"change_registry_num,omitempty"`

	// 修改数量
	ModifyNum *int32 `json:"modify_num,omitempty"`

	// 新增数量
	AddNum *int32 `json:"add_num,omitempty"`

	// 删除数量
	DeleteNum *int32 `json:"delete_num,omitempty"`

	// trust num
	TrustNum *int32 `json:"trust_num,omitempty"`

	// untrust num
	UntrustNum *int32 `json:"untrust_num,omitempty"`

	// unknown_num
	UnknownNum     *int32 `json:"unknown_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFileStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowFileStatisticResponse", string(data)}, " ")
}
