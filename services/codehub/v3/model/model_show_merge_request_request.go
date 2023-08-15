package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMergeRequestRequest Request Object
type ShowMergeRequestRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	// 合并请求id
	MergeRequestId int32 `json:"merge_request_id"`
}

func (o ShowMergeRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestRequest struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestRequest", string(data)}, " ")
}
