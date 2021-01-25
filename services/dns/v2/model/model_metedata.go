package model

import (
	"encoding/json"

	"strings"
)

type Metedata struct {
	// 满足查询条件的资源总数，不受分页（即limit、offset参数）影响。
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o Metedata) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Metedata struct{}"
	}

	return strings.Join([]string{"Metedata", string(data)}, " ")
}
