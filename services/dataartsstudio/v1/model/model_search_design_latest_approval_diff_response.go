package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDesignLatestApprovalDiffResponse Response Object
type SearchDesignLatestApprovalDiffResponse struct {
	Data           *PublishVersionVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o SearchDesignLatestApprovalDiffResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDesignLatestApprovalDiffResponse struct{}"
	}

	return strings.Join([]string{"SearchDesignLatestApprovalDiffResponse", string(data)}, " ")
}
