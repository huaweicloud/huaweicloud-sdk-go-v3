package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiAppsResponse Response Object
type BatchDeleteAiAppsResponse struct {

	// 已删除的AI应用ID列表
	AppIds         *[]string `json:"appIds,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAiAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiAppsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiAppsResponse", string(data)}, " ")
}
