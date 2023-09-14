package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVolumeQoSResponse Response Object
type ModifyVolumeQoSResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyVolumeQoSResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeQoSResponse struct{}"
	}

	return strings.Join([]string{"ModifyVolumeQoSResponse", string(data)}, " ")
}
