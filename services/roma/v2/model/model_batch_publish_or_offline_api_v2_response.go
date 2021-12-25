package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchPublishOrOfflineApiV2Response struct {
	// 发布或下线成功的信息

	Success *[]PublishResp `json:"success,omitempty"`
	// 发布或下线失败的API及错误信息

	Failure        *[]BatchFailure `json:"failure,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchPublishOrOfflineApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublishOrOfflineApiV2Response struct{}"
	}

	return strings.Join([]string{"BatchPublishOrOfflineApiV2Response", string(data)}, " ")
}
