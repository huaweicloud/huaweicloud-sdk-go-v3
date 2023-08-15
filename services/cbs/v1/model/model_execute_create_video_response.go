package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCreateVideoResponse Response Object
type ExecuteCreateVideoResponse struct {

	// 视频id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCreateVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCreateVideoResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCreateVideoResponse", string(data)}, " ")
}
