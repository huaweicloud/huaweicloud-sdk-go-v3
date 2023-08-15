package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHostListRequestBody 查询主机信息请求体
type GetHostListRequestBody struct {

	// 主机ID列表。可以根据主机ID列表进行批量过滤
	HostIdList *[]string `json:"host_id_list,omitempty"`

	Filter *GetHostListFilter `json:"filter"`
}

func (o GetHostListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostListRequestBody struct{}"
	}

	return strings.Join([]string{"GetHostListRequestBody", string(data)}, " ")
}
