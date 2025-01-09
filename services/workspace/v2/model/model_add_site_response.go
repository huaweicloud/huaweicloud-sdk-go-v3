package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSiteResponse Response Object
type AddSiteResponse struct {

	// 初始化站点的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSiteResponse struct{}"
	}

	return strings.Join([]string{"AddSiteResponse", string(data)}, " ")
}
