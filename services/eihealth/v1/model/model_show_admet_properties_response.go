package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetPropertiesResponse Response Object
type ShowAdmetPropertiesResponse struct {

	// 分子ADMET属性字典
	Body           map[string]interface{} `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAdmetPropertiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetPropertiesResponse struct{}"
	}

	return strings.Join([]string{"ShowAdmetPropertiesResponse", string(data)}, " ")
}
