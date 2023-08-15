package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryByCloudIdeResponse Response Object
type ShowRepositoryByCloudIdeResponse struct {

	// CloudIde打开链接:https://xxx/cloudide/loading?instanceId=xxx&scmUrl=xxx
	Url            *string `json:"url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryByCloudIdeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryByCloudIdeResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryByCloudIdeResponse", string(data)}, " ")
}
