package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleDdmsResponse Response Object
type ShowAvalibleDdmsResponse struct {

	// 可用目标DDM。
	Instances      *[]DdmInstance4Restore `json:"instances,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAvalibleDdmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleDdmsResponse struct{}"
	}

	return strings.Join([]string{"ShowAvalibleDdmsResponse", string(data)}, " ")
}
