package model

import (
	"encoding/json"

	"strings"
)

// 截图信息。仅当截图成功后才能查询到此信息，未截图、正在截图以及截图失败时，此字段为空。
type ThumbnailInfo struct {
	Sample *[]ThumbnailRsp `json:"sample,omitempty"`

	Dots *[]ThumbnailRsp `json:"dots,omitempty"`

	ExecDesc *string `json:"exec_desc,omitempty"`

	ThumbnailStatus *string `json:"thumbnail_status,omitempty"`
}

func (o ThumbnailInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ThumbnailInfo struct{}"
	}

	return strings.Join([]string{"ThumbnailInfo", string(data)}, " ")
}
