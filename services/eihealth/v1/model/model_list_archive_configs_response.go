package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArchiveConfigsResponse Response Object
type ListArchiveConfigsResponse struct {

	// 归档设置记录总数
	Count *int32 `json:"count,omitempty"`

	// 配置项
	Configs        *[]GetArchiveConfigRsp `json:"configs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListArchiveConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArchiveConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListArchiveConfigsResponse", string(data)}, " ")
}
