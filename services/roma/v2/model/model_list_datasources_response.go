package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDatasourcesResponse struct {

	// 返回所有满足条件的对象个数
	Total *int64 `json:"total,omitempty" xml:"total"`

	// 返回对象的大小
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 返回的实体对象
	Entities       *[]DataSourceRsp `json:"entities,omitempty" xml:"entities"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDatasourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDatasourcesResponse", string(data)}, " ")
}
