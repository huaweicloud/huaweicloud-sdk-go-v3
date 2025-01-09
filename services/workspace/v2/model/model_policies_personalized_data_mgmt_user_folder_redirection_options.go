package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions 用户文件夹重定向选项。
type PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions struct {

	// 还原本地目录启用
	RestoreLocalDirectoriesEnable *bool `json:"restore_local_directories_enable,omitempty"`

	// 应用数据漫游启用
	AppDataRoamingEnale *bool `json:"appData_roaming_enale,omitempty"`

	// 桌面重定向
	RedirecDesktop *bool `json:"redirec_desktop,omitempty"`

	// 开始菜单
	RedirecStartMenu *bool `json:"redirec_start_menu,omitempty"`

	// 文档
	RedirecDocuments *bool `json:"redirec_documents,omitempty"`

	// 照片
	RedirecPictures *bool `json:"redirec_pictures,omitempty"`

	// 音乐
	RedirecMusic *bool `json:"redirec_music,omitempty"`

	// 录音
	RedirecVideos *bool `json:"redirec_videos,omitempty"`

	// 最爱
	RedirecFavorites *bool `json:"redirec_favorites,omitempty"`

	// 录音
	RedirecContacts *bool `json:"redirec_contacts,omitempty"`

	// 下载
	RedirecDownloads *bool `json:"redirec_downloads,omitempty"`

	// 链接
	RedirecLinks *bool `json:"redirec_links,omitempty"`

	// 查找
	RedirecSearches *bool `json:"redirec_searches,omitempty"`

	// 游戏存储
	RedirecSavedGames *bool `json:"redirec_saved_games,omitempty"`
}

func (o PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions", string(data)}, " ")
}
