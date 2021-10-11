package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddingBackendInstancesV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询到的云服务器列表

	Members        *[]VpcMemberInfo `json:"members,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o AddingBackendInstancesV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddingBackendInstancesV2Response struct{}"
	}

	return strings.Join([]string{"AddingBackendInstancesV2Response", string(data)}, " ")
}
