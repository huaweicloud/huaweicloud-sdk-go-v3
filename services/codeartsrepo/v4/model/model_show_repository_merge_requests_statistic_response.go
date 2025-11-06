package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryMergeRequestsStatisticResponse Response Object
type ShowRepositoryMergeRequestsStatisticResponse struct {
	Body           *[]MergeRequestStatisticDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowRepositoryMergeRequestsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryMergeRequestsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryMergeRequestsStatisticResponse", string(data)}, " ")
}
