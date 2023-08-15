package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviewSettingRequest Request Object
type ShowReviewSettingRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`
}

func (o ShowReviewSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviewSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowReviewSettingRequest", string(data)}, " ")
}
