package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFaceSetResponse Response Object
type ShowFaceSetResponse struct {
	FaceSetInfo    *FaceSetInfo `json:"face_set_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowFaceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFaceSetResponse struct{}"
	}

	return strings.Join([]string{"ShowFaceSetResponse", string(data)}, " ")
}
