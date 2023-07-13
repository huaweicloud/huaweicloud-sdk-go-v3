package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApproveOpinionDetail Approve Detail
type ApproveOpinionDetail struct {

	// Approve Result.
	Result *string `json:"result,omitempty"`

	// Approve content.
	Content *string `json:"content,omitempty"`
}

func (o ApproveOpinionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproveOpinionDetail struct{}"
	}

	return strings.Join([]string{"ApproveOpinionDetail", string(data)}, " ")
}
