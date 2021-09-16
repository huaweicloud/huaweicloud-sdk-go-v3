package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateDatabaseWaterMarkResponse struct {
	// 嵌入水印后的数据

	MarkedData     *[]map[string]interface{} `json:"marked_data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o CreateDatabaseWaterMarkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseWaterMarkResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseWaterMarkResponse", string(data)}, " ")
}
