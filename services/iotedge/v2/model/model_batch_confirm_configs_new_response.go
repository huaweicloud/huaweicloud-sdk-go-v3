package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchConfirmConfigsNewResponse Response Object
type BatchConfirmConfigsNewResponse struct {

	// 已确认的配置项id
	Ids            *interface{} `json:"ids,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchConfirmConfigsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmConfigsNewResponse struct{}"
	}

	return strings.Join([]string{"BatchConfirmConfigsNewResponse", string(data)}, " ")
}
