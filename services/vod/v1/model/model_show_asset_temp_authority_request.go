package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAssetTempAuthorityRequest struct {
	// http方法<br/>初始化多段上传任务为POST,列举多段上传任务为GET,<br/> 上传段为PUT,合并段为POST,列举已上传段为GET,取消段合并为DELETE。

	HttpVerb string `json:"http_verb"`
	// 上传段时,每段媒资内容的md5值,非上传段操作不涉及此字段<br/>

	ContentMd5 *string `json:"content_md5,omitempty"`
	// 上传段时,媒资内容对应的content-type值,非上传段操作不涉及此字段<br/>

	ContentType *string `json:"content_type,omitempty"`
	// 桶名<br/>

	Bucket string `json:"bucket"`
	// 对象名<br/>

	ObjectKey string `json:"object_key"`
	// 上传任务id,由OBS分配<br/>

	UploadId *string `json:"upload_id,omitempty"`
	// 上传段号,取值[1,10000]<br/>

	PartNumber *int32 `json:"part_number,omitempty"`
}

func (o ShowAssetTempAuthorityRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAssetTempAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetTempAuthorityRequest", string(data)}, " ")
}
