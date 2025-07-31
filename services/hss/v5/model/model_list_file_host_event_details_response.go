package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileHostEventDetailsResponse Response Object
type ListFileHostEventDetailsResponse struct {

	// 文件总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 变更总数
	ChangeTotalNum *int32 `json:"change_total_num,omitempty"`

	// 变更文件数量
	ChangeFileNum *int32 `json:"change_file_num,omitempty"`

	// 变更注册表数量
	ChangeRegistryNum *int32 `json:"change_registry_num,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 变更文件信息列表
	DataList       *[]FileHostEventDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListFileHostEventDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileHostEventDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListFileHostEventDetailsResponse", string(data)}, " ")
}
