package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadDataRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 目标文件夹
	TargetFolder *string `json:"target_folder,omitempty"`

	// 分段序号，表示第几个文件片段
	PartNumber *int32 `json:"part_number,omitempty"`

	// 分段总数，上传的文件总共分成了几个片段
	TotalPart *int32 `json:"total_part,omitempty"`

	// 分段上传任务id，除了第一个片段外，后续的片段都需要标识出任务id
	MultipartId *string `json:"multipart_id,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// MD5
	Md5 *string `json:"md5,omitempty"`

	Body *UploadDataRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDataRequest struct{}"
	}

	return strings.Join([]string{"UploadDataRequest", string(data)}, " ")
}
