package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesLevelTagsRequest Request Object
type ListResourcesLevelTagsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`
}

func (o ListResourcesLevelTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesLevelTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesLevelTagsRequest", string(data)}, " ")
}
