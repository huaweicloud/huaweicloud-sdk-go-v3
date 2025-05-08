package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRealTimeClipRequestBody obs桶取回请求消息
type CreateRealTimeClipRequestBody struct {

	// 时移域名
	TimeshiftDomain *string `json:"timeshift_domain,omitempty"`

	// live的应用名，默认可填写live
	App *string `json:"app,omitempty"`

	// 截取的流名
	Stream string `json:"stream"`

	// 直播时移的转码模板，空表示截取源流
	TransTemplateName *string `json:"trans_template_name,omitempty"`

	// 截取流的开始时间，格式为：YYYY-MM-DDTHH:mm:ssZ（UTC时间）,开始时间与结束时间的间隔最大为12小时
	StartTime string `json:"start_time"`

	// 截取流的开始时间，格式为：YYYY-MM-DDTHH:mm:ssZ（UTC时间）,开始时间与结束时间的间隔最大为12小时
	EndTime string `json:"end_time"`

	// 0 非固化，1 固化。默认非固化
	IsPersistence *int32 `json:"is_persistence,omitempty"`

	// 非固化场景为空，固化场景必填，复制到该桶中。
	OutputBucket *string `json:"output_bucket,omitempty"`

	// 截取指定的m3u8文件路径名，仅支持HLS
	OutputObject *string `json:"output_object,omitempty"`

	MediaProcessTask *AdditionalObjectProcessReq `json:"media_process_task,omitempty"`

	// 回调地址，为空则不回调
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文，回调时会回调给用户，透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o CreateRealTimeClipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRealTimeClipRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRealTimeClipRequestBody", string(data)}, " ")
}
