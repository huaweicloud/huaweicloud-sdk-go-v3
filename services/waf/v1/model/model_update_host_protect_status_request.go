package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateHostProtectStatusRequest struct {
	// 域名Id（通过查询云模式防护域名列表获取域名id)

	InstanceId string `json:"instance_id"`

	Body *UpdateHostProtectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateHostProtectStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusRequest", string(data)}, " ")
}
