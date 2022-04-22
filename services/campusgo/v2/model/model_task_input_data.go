package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输入数据的配置信息
type TaskInputData struct {

	// VIS的视频流名称，当输入为vis类型时必选
	StreamName *string `json:"stream_name,omitempty"`

	// OBS桶名，当输入为obs类型是必选
	Bucket *string `json:"bucket,omitempty"`

	// OBS的路径，当输入为obs类型时必选
	Path *string `json:"path,omitempty"`

	// url输入源的地址或者获取视频流地址的restful请求地址，当输入为url或者edgerestful类型时必选
	Url *string `json:"url,omitempty"`

	// 获取视频流的restful请求携带的请求头，当输入为edgerestful类型时可选
	Headers *interface{} `json:"headers,omitempty"`

	// 是否需要对https请求进行证书校验，当输入为edgerestful类型时必选
	CertificateCheck *bool `json:"certificate_check,omitempty"`

	// restful请求返回的body中，视频流地址的路径，当输入为edgerestful类型时必选
	RtspPathInResponse *string `json:"rtsp_path_in_response,omitempty"`

	// VCN设备ID，当输入为vcn类型时必选
	DeviceId *string `json:"device_id,omitempty"`

	// 准备进行分析的码流，其中1代表主码流，2代表子码流1,3代表子码流2，当输入为vcn类型时可选
	StreamType *int32 `json:"stream_type,omitempty"`

	// IEF挂载的边缘设备的ID，当输入为edgecamera类型时必选
	Id *string `json:"id,omitempty"`

	// 可选，当前输入的序号，从0开始递增，不可重复
	Index *int32 `json:"index,omitempty"`
}

func (o TaskInputData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInputData struct{}"
	}

	return strings.Join([]string{"TaskInputData", string(data)}, " ")
}
