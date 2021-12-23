package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAllFaceSetsResponse struct {
	// 人脸库信息集合，详见[FaceSetInfo](https://support.huaweicloud.com/api-face/face_02_0020.html)。 调用失败时无此字段。

	FaceSetsInfo   *[]FaceSetInfo `json:"face_sets_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowAllFaceSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllFaceSetsResponse struct{}"
	}

	return strings.Join([]string{"ShowAllFaceSetsResponse", string(data)}, " ")
}
