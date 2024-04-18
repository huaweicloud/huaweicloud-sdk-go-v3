package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFileResponse Response Object
type CreateFileResponse struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件上传地址，有效期为24小时。 > * [调用OBS的[“PUT上传”](https://support.huaweicloud.com/api-obs/obs_04_0080.html)接口上传文件。](tag:hc) > * [调用OBS的[“PUT上传”](https://support.huaweicloud.com/intl/zh-cn/api-obs/obs_04_0080.html)接口上传文件。](tag:hk) > * [调用OBS的“PUT上传”接口上传文件。](tag:cmcc) > * 调用上述接口时，Content-MD5头必须填写，填写的值跟file_md5中的值相同。
	UploadUrl *string `json:"upload_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFileResponse struct{}"
	}

	return strings.Join([]string{"CreateFileResponse", string(data)}, " ")
}
