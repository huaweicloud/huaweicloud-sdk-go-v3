package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchConfirmConfigsResponse Response Object
type BatchConfirmConfigsResponse struct {

	// 已确认的配置项id
	Ids            *interface{} `json:"ids,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchConfirmConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmConfigsResponse struct{}"
	}

	return strings.Join([]string{"BatchConfirmConfigsResponse", string(data)}, " ")
}
