package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLtsLogDetailsResponse struct {

	// 日志组名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 日志组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 日志流id
	StreamId *string `json:"stream_id,omitempty" xml:"stream_id"`

	// 日志流名称
	StreamName     *string `json:"stream_name,omitempty" xml:"stream_name"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLtsLogDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsLogDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowLtsLogDetailsResponse", string(data)}, " ")
}
