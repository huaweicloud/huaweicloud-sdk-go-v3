package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasourcesResponse Response Object
type ListDatasourcesResponse struct {

	// 返回所有满足条件的对象个数
	Total *int64 `json:"total,omitempty"`

	// 返回对象的大小
	Size *int32 `json:"size,omitempty"`

	// 返回的实体对象
	Entities       *[]DataSourceRsp `json:"entities,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDatasourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDatasourcesResponse", string(data)}, " ")
}
