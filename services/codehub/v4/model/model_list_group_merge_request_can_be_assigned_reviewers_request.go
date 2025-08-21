package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMergeRequestCanBeAssignedReviewersRequest Request Object
type ListGroupMergeRequestCanBeAssignedReviewersRequest struct {

	// **参数解释：** 代码组id，代码组首页，Group ID后的数字Id
	GroupId int32 `json:"group_id"`
}

func (o ListGroupMergeRequestCanBeAssignedReviewersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMergeRequestCanBeAssignedReviewersRequest struct{}"
	}

	return strings.Join([]string{"ListGroupMergeRequestCanBeAssignedReviewersRequest", string(data)}, " ")
}
