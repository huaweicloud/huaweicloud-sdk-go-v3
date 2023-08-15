package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandVolumesResponse Response Object
type ExpandVolumesResponse struct {

	// 扩容磁盘任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumesResponse struct{}"
	}

	return strings.Join([]string{"ExpandVolumesResponse", string(data)}, " ")
}
