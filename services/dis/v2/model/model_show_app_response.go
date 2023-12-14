package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppResponse Response Object
type ShowAppResponse struct {

	// App的名称。
	AppName *string `json:"app_name,omitempty"`

	// App的唯一标识符。
	AppId *string `json:"app_id,omitempty"`

	// App创建的时间，单位毫秒。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 关联通道列表。
	CommitCheckpointStreamNames *[]string `json:"commit_checkpoint_stream_names,omitempty"`
	HttpStatusCode              int       `json:"-"`
}

func (o ShowAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResponse struct{}"
	}

	return strings.Join([]string{"ShowAppResponse", string(data)}, " ")
}
