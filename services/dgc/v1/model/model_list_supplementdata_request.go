package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupplementdataRequest Request Object
type ListSupplementdataRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`
}

func (o ListSupplementdataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplementdataRequest struct{}"
	}

	return strings.Join([]string{"ListSupplementdataRequest", string(data)}, " ")
}
