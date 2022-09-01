package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteFaceSetResponse struct {

	// 人脸库名称。 调用失败时无此字段。
	FaceSetName    *string `json:"face_set_name,omitempty" xml:"face_set_name"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFaceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFaceSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteFaceSetResponse", string(data)}, " ")
}
