package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDatabaseDataResponse struct {

	// 总条数
	Count *int32 `json:"count,omitempty"`

	// 查询记录列表
	Objects        *[]map[string]interface{} `json:"objects,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDatabaseDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseDataResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseDataResponse", string(data)}, " ")
}
