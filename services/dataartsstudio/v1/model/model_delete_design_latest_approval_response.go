package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignLatestApprovalResponse Response Object
type DeleteDesignLatestApprovalResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignLatestApprovalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignLatestApprovalResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignLatestApprovalResponse", string(data)}, " ")
}
