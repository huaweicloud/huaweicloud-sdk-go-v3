package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFaceSetResponse struct {
	FaceSetInfo    *FaceSetInfo `json:"face_set_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateFaceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFaceSetResponse struct{}"
	}

	return strings.Join([]string{"CreateFaceSetResponse", string(data)}, " ")
}
