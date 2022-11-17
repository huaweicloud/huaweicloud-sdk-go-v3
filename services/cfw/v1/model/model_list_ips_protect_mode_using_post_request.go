package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIpsProtectModeUsingPostRequest struct {

	// 防护对象id
	ObjectId string `json:"object_id"`
}

func (o ListIpsProtectModeUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsProtectModeUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ListIpsProtectModeUsingPostRequest", string(data)}, " ")
}
