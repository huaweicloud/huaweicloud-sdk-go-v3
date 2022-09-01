package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFaceSetRequest struct {
	Body *CreateFaceSetReq `json:"body,omitempty" xml:"body"`
}

func (o CreateFaceSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFaceSetRequest struct{}"
	}

	return strings.Join([]string{"CreateFaceSetRequest", string(data)}, " ")
}
