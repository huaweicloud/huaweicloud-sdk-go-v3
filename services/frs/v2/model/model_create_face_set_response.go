package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFaceSetResponse Response Object
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
