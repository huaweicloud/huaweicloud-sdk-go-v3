package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApplicationRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 资源空间ID，唯一标识一个资源空间，由物联网平台在创建资源空间时分配。资源空间对应的是物联网平台原有的应用，在物联网平台的含义与应用一致，只是变更了名称。

	AppId string `json:"app_id"`
}

func (o ShowApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowApplicationRequest", string(data)}, " ")
}
