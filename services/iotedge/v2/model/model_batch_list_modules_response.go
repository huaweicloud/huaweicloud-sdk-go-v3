package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchListModulesResponse struct {
	// 总记录数

	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`
	// 每页记录数

	Modules        *[]EdgeModuleRespDto `json:"modules,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchListModulesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListModulesResponse struct{}"
	}

	return strings.Join([]string{"BatchListModulesResponse", string(data)}, " ")
}
