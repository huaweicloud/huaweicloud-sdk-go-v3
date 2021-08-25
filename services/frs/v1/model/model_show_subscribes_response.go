package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSubscribesResponse struct {
	// 调用成功时表示最大的人脸库数量。 调用失败时无此字段。

	MaxFaceSetNumber *int32 `json:"max_face_set_number,omitempty"`

	DetectService *ServiceInfo `json:"detect_service,omitempty"`

	LiveDetectService *ServiceInfo `json:"live_detect_service,omitempty"`

	CompareService *ServiceInfo `json:"compare_service,omitempty"`

	SearchService  *ServiceInfo `json:"search_service,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowSubscribesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubscribesResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscribesResponse", string(data)}, " ")
}
