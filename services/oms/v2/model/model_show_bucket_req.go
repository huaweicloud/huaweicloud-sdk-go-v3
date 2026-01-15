package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBucketReq 查询桶请求体
type ShowBucketReq struct {

	// 云类型 AWS：亚马逊 Aliyun：阿里云 Qiniu：七牛云 QingCloud：青云 Tencent：腾讯云 Baidu：百度云 KingsoftCloud：金山云 Azure：微软云 UCloud：优刻得 HuaweiCloud：华为云 Google: 谷歌云 URLSource：URL HEC：HEC
	CloudType string `json:"cloud_type"`

	// 目标桶中需要查询的对象文件路径，/结尾
	FilePath string `json:"file_path"`

	// 源端桶的AK（最大长度100个字符）
	Ak *string `json:"ak,omitempty"`

	// 源端桶的SK（最大长度100个字符）
	Sk *string `json:"sk,omitempty"`

	// 连接字符串，用于微软云Blob鉴权
	ConnectionString *string `json:"connection_string,omitempty"`

	// 用于谷歌云Cloud Storage鉴权
	JsonAuthFile *string `json:"json_auth_file,omitempty"`

	// 数据中心，区域
	DataCenter string `json:"data_center"`

	// 分页信息，页大小
	PageSize int32 `json:"page_size"`

	// 分页信息，当前页最后一个对象名称（偏移量）
	BehindFilename string `json:"behind_filename"`

	// 腾讯云APPID，当源端为腾讯云时，需要填写此参数，您可以在腾讯云控制台账号信息页面获取。
	AppId *string `json:"app_id,omitempty"`

	// 桶名
	BucketName string `json:"bucket_name"`
}

func (o ShowBucketReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBucketReq struct{}"
	}

	return strings.Join([]string{"ShowBucketReq", string(data)}, " ")
}
