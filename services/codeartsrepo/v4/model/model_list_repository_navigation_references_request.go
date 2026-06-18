package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryNavigationReferencesRequest Request Object
type ListRepositoryNavigationReferencesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/intl/zh-cn/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk_ch)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过100000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 版本提交id **取值范围：** 不涉及
	Revision *string `json:"revision,omitempty"`

	// **参数解释：** 引用，可以是分支名称、标签名称或者commitid，如果不传则为默认分支。 **取值范围：** 字符串长度不少于1，不超过2000。
	Ref *string `json:"ref,omitempty"`

	// **参数解释：** 搜索符号（页面选中的字符串） **取值范围：** 不涉及
	Symbol string `json:"symbol"`

	// **参数解释：** 代码语言 **取值范围：** - C - C++ - Go - Java - JavaScript - PHP - Python - Ruby - Rust
	Language string `json:"language"`

	// **参数解释：** blob文件ID。通过[[查询某个仓库的文件信息](https://support.huaweicloud.com/api-codeartsrepo/ListFilesByQuery.html)](tag:hws)[[查询某个仓库的文件信息](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListFilesByQuery.html)](tag:hws_hk)[查询某个仓库的文件信息](tag:hcs,hcs_sm)接口查询某个仓库的文件信息获取。 **取值范围：** 不涉及。
	Blob string `json:"blob"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过10000。
	FilePath string `json:"file_path"`
}

func (o ListRepositoryNavigationReferencesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryNavigationReferencesRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryNavigationReferencesRequest", string(data)}, " ")
}
