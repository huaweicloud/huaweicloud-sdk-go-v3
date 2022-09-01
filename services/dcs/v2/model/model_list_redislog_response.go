package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRedislogResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty" xml:"total_num"`

	// 运行日志列表
	FileList       *[]RunlogItem `json:"file_list,omitempty" xml:"file_list"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRedislogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedislogResponse struct{}"
	}

	return strings.Join([]string{"ListRedislogResponse", string(data)}, " ")
}
