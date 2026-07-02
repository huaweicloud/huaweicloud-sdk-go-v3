package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventDataResponse Response Object
type ShowEventDataResponse struct {

	// **参数解释**： 配置信息列表 **取值范围**： 不涉及
	Datapoints     *[]EventDataInfo `json:"datapoints,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowEventDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDataResponse struct{}"
	}

	return strings.Join([]string{"ShowEventDataResponse", string(data)}, " ")
}
