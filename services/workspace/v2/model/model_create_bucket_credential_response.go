package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBucketCredentialResponse Response Object
type CreateBucketCredentialResponse struct {

	// 访问的服务终端节点。
	ServerEndPoint *string `json:"server_end_point,omitempty"`

	// 存放的桶名称。
	BucketName *string `json:"bucket_name,omitempty"`

	// OBS对象路径。 注: path是对象在obs中的完整路径。 例如桶存在如下目录结构的数据。   Bucket:     ├─dir1     | ├─object1.txt     | └─object2.txt     └─object3.txt Object1的path: dir1/object1.txt Object2的path: dir1/object2.txt Object3的path: object3.txt
	ObjectPath *string `json:"object_path,omitempty"`

	Policy *ObsPolicy `json:"policy,omitempty"`

	Credential     *ObsCredential `json:"credential,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateBucketCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBucketCredentialResponse struct{}"
	}

	return strings.Join([]string{"CreateBucketCredentialResponse", string(data)}, " ")
}
