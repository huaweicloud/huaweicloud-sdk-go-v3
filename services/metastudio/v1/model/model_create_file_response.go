package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFileResponse Response Object
type CreateFileResponse struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件上传地址，有效期为24小时。 > * [调用OBS的[“PUT上传”](https://support.huaweicloud.com/api-obs/obs_04_0080.html)接口上传文件。](tag:hc) > * [调用OBS的[“PUT上传”](https://support.huaweicloud.com/intl/zh-cn/api-obs/obs_04_0080.html)接口上传文件。](tag:hk) > * [调用OBS的“PUT上传”接口上传文件。](tag:cmcc) > * 调用上述接口时，Content-MD5头必须填写，填写的值跟file_md5中的值相同，md5值获取详情请参考[使用Java代码生成文件内容的MD5值](metastudio_02_0052.xml)。 > * 调用上述接口时，Content-Type头必须填写，填写的值根据不同的文件类型有所不同。     文件类型为gif，Content-Type填写image/gif     文件类型为jpeg、jpg，Content-Type填写image/jpeg     文件类型为png，Content-Type填写image/png     文件类型为mp4，Content-Type填写video/mp4     文件类型为mp3，Content-Type填写audio/mp3     文件类型为wav，Content-Type填写audio/wav     其余所有类型，Content-Type填写application/octet-stream
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
