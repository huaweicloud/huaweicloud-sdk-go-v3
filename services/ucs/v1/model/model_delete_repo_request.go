package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoRequest Request Object
type DeleteRepoRequest struct {

	// 仓库id
	Repoid string `json:"repoid"`
}

func (o DeleteRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoRequest struct{}"
	}

	return strings.Join([]string{"DeleteRepoRequest", string(data)}, " ")
}
