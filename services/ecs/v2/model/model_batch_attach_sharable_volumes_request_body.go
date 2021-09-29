package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type BatchAttachSharableVolumesRequestBody struct {
	// 共享磁盘需要挂载的弹性云服务器列表。

	Serverinfo []BatchAttachSharableVolumesOption `json:"serverinfo"`
}

func (o BatchAttachSharableVolumesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequestBody", string(data)}, " ")
}
