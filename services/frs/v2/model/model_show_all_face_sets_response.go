package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAllFaceSetsResponse struct {
	// 人脸库信息集合，详见[FaceSetInfo](zh-cn_topic_0106912072.xml)。 调用失败时无此字段。

	FaceSetsInfo   *[]FaceSetInfo `json:"face_sets_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAllFaceSetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAllFaceSetsResponse struct{}"
	}

	return strings.Join([]string{"ShowAllFaceSetsResponse", string(data)}, " ")
}
