package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlInstanceListResponse struct {
	// 实例列表信息。

	Instances *[]MysqlInstanceListInfo `json:"instances,omitempty"`
	// 总记录数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMysqlInstanceListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlInstanceListResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlInstanceListResponse", string(data)}, " ")
}
