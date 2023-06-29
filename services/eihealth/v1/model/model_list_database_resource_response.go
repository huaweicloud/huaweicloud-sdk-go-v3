package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseResourceResponse Response Object
type ListDatabaseResourceResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 数据库资源列表
	Resources *[]DatabaseResourceRsp `json:"resources,omitempty"`

	XResourceMappings *string `json:"X-Resource-Mappings,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ListDatabaseResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseResourceResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseResourceResponse", string(data)}, " ")
}
