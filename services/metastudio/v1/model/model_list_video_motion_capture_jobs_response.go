package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVideoMotionCaptureJobsResponse Response Object
type ListVideoMotionCaptureJobsResponse struct {

	// 视频驱动任务总数。
	Total *int32 `json:"total,omitempty"`

	// 视频驱动任务列表。
	VideoMotionCaptureJobs *[]VideoMotionCaptureInfo `json:"video_motion_capture_jobs,omitempty"`
	HttpStatusCode         int                       `json:"-"`
}

func (o ListVideoMotionCaptureJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVideoMotionCaptureJobsResponse struct{}"
	}

	return strings.Join([]string{"ListVideoMotionCaptureJobsResponse", string(data)}, " ")
}
