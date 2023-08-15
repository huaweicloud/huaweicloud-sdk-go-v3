package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyPortResponse Response Object
type UpdateProxyPortResponse struct {

	// 修改proxy端口号的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxyPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyPortResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyPortResponse", string(data)}, " ")
}
