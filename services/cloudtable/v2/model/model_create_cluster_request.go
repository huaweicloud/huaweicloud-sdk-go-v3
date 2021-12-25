package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateClusterRequest struct {
	// 语言类型

	XLanguage string `json:"X-Language"`

	Body *CreateClusterRequestBody `json:"body,omitempty"`
}

func (o CreateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterRequest", string(data)}, " ")
}
