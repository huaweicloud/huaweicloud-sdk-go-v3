package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlobsRequest Request Object
type ShowBlobsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** blob文件ID。通过[[查询某个仓库的文件信息](https://support.huaweicloud.com/api-codeartsrepo/ListFilesByQuery.html)](tag:hws)[[查询某个仓库的文件信息](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListFilesByQuery.html)](tag:hws_hk)[[查询某个仓库的文件信息](https://support.huaweicloud.com/eu/api-codeartsrepo/ListFilesByQuery.html)](tag:hws_eu)[查询某个仓库的文件信息](tag:hcs,hcs_sm)接口查询某个仓库的文件信息获取。 **取值范围：** 不涉及。
	BlobId string `json:"blob_id"`
}

func (o ShowBlobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlobsRequest struct{}"
	}

	return strings.Join([]string{"ShowBlobsRequest", string(data)}, " ")
}
