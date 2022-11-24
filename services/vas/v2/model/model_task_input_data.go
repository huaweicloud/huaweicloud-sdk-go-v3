package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输入数据的配置信息
type TaskInputData struct {

	// VIS的视频流名称，当输入为vis类型时必填。
	StreamName *string `json:"stream_name,omitempty"`

	// OBS桶名，当输入为obs类型时必填。
	Bucket *string `json:"bucket,omitempty"`

	// OBS的路径，当输入为obs类型时必填。
	Path *string `json:"path,omitempty"`

	// url输入源的地址或者获取视频流地址的restful请求地址，当输入为url类型或者edgerestful类型时必填。长度不超过1000。
	Url *string `json:"url,omitempty"`

	// 获取视频流地址的restful请求携带的请求头，当输入为edgerestful类型时可选。整体呈json格式，以键值对的形式表示请求头和取值，最多允许10组。
	Headers *interface{} `json:"headers,omitempty"`

	// 是否需要对https请求进行证书校验，当输入为edgerestful类型时必填。取值为true或者false。
	CertificateCheck *bool `json:"certificate_check,omitempty"`

	// restful请求返回的body中，视频流地址的路径，当输入为edgerestful类型时必填。长度不超过1024。
	RtspPathInResponse *string `json:"rtsp_path_in_response,omitempty"`

	// IEF节点的ID，仅部分服务在输入类型为edgerestful或vcn时需填且必填。
	NodeId *string `json:"node_id,omitempty"`

	// VCN设备ID，当输入为vcn类型时必填。
	DeviceId *string `json:"device_id,omitempty"`

	// 准备进行分析的码流，当输入为vcn类型时选填。取值范围为1~3，其中1代表主码流，2代表子码流1,3代表子码流2。
	StreamType *int32 `json:"stream_type,omitempty"`

	// IEF挂载的边缘设备的ID，当输入为edgecamera类型时必填。
	Id *string `json:"id,omitempty"`
}

func (o TaskInputData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInputData struct{}"
	}

	return strings.Join([]string{"TaskInputData", string(data)}, " ")
}
