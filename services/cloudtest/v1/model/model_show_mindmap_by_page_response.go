package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapByPageResponse Response Object
type ShowMindmapByPageResponse struct {
	Params         *MindmapPageParamV3 `json:"params,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowMindmapByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowMindmapByPageResponse", string(data)}, " ")
}
