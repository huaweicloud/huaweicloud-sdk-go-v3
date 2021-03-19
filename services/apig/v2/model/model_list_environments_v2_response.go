package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEnvironmentsV2Response struct {
	// 符合条件的环境总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回的列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次返回的环境列表

	Envs           *[]EnvResp `json:"envs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListEnvironmentsV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnvironmentsV2Response struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsV2Response", string(data)}, " ")
}
