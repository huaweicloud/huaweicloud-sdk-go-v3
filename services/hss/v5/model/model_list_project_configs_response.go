package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectConfigsResponse Response Object
type ListProjectConfigsResponse struct {

	// 配置信息列表
	DataList       *[]ProjectConfigInfo `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListProjectConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectConfigsResponse", string(data)}, " ")
}
