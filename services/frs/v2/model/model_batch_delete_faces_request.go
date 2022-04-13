package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteFacesRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *DeleteFacesBatchReq `json:"body,omitempty"`
}

func (o BatchDeleteFacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFacesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFacesRequest", string(data)}, " ")
}
