package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSupplementdataRequest Request Object
type CreateSupplementdataRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *SupplementDataInfo `json:"body,omitempty"`
}

func (o CreateSupplementdataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSupplementdataRequest struct{}"
	}

	return strings.Join([]string{"CreateSupplementdataRequest", string(data)}, " ")
}
