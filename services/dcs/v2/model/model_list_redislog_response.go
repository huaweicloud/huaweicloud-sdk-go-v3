package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRedislogResponse struct {
	// 总数

	TotalNum *int32 `json:"total_num,omitempty"`
	// 运行日志列表

	FileList       *[]RunlogItem `json:"file_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListRedislogResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRedislogResponse struct{}"
	}

	return strings.Join([]string{"ListRedislogResponse", string(data)}, " ")
}
