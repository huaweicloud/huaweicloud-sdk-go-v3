package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 截图信息。  > 仅当截图成功后才能查询到此信息，未截图、正在截图以及截图失败时，无此字段信息。
type ThumbnailInfo struct {

	// 视频截图信息。
	Sample *[]ThumbnailRsp `json:"sample,omitempty" xml:"sample"`

	Dots *[]ThumbnailRsp `json:"dots,omitempty" xml:"dots"`

	// 执行情况描述。
	ExecDesc *string `json:"exec_desc,omitempty" xml:"exec_desc"`

	// 截图状态。  取值如下： - UN_THUMBNAIL：未截图 - THUMBNAILING：截图中 - THUMBNAIL_SUCCEED：截图成功 - THUMBNAIL_FAILED：截图失败
	ThumbnailStatus *string `json:"thumbnail_status,omitempty" xml:"thumbnail_status"`
}

func (o ThumbnailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbnailInfo struct{}"
	}

	return strings.Join([]string{"ThumbnailInfo", string(data)}, " ")
}
