package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKernelModuleHostInfoResponse Response Object
type ListKernelModuleHostInfoResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 服务器列表
	DataList       *[]KernelModuleHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListKernelModuleHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKernelModuleHostInfoResponse struct{}"
	}

	return strings.Join([]string{"ListKernelModuleHostInfoResponse", string(data)}, " ")
}
