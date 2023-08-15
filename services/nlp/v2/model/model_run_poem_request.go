package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPoemRequest Request Object
type RunPoemRequest struct {
	Body *CreatePoem `json:"body,omitempty"`
}

func (o RunPoemRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPoemRequest struct{}"
	}

	return strings.Join([]string{"RunPoemRequest", string(data)}, " ")
}
