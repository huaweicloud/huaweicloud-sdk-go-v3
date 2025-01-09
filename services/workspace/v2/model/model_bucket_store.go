package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BucketStore OBS桶存储。
type BucketStore struct {

	// 桶名称固定格式:wks-appcenter-{project_id}; 需先调用桶授权接口进行授权。
	BucketName *string `json:"bucket_name,omitempty"`

	// OBS对象路径。 注: bucket_file_path是对象在obs中的完整路径,不能以/开头。 例如桶存在如下目录结构的数据。   Bucket:     ├─dir1     | ├─object1.txt     | └─object2.txt     └─object3.txt Object1的路径: dir1/object1.txt Object2的路径: dir1/object2.txt Object3的路径: object3.txt
	BucketFilePath *string `json:"bucket_file_path,omitempty"`
}

func (o BucketStore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketStore struct{}"
	}

	return strings.Join([]string{"BucketStore", string(data)}, " ")
}
