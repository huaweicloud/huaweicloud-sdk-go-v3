package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FolderRedirectionV2Options 文件夹重定向配置(v2)。
type FolderRedirectionV2Options struct {
	AppDataRoamingConfigurations *AppDataRoamingConfigurations `json:"app_data_roaming_configurations,omitempty"`

	DesktopConfigurations *DesktopConfigurations `json:"desktop_configurations,omitempty"`

	StartMenuConfigurations *StartMenuConfigurations `json:"start_menu_configurations,omitempty"`

	DocumentsConfigurations *DocumentsConfigurations `json:"documents_configurations,omitempty"`

	PicturesConfigurations *PicturesConfigurations `json:"pictures_configurations,omitempty"`

	MusicConfigurations *MusicConfigurations `json:"music_configurations,omitempty"`

	VideosConfigurations *VideosConfigurations `json:"videos_configurations,omitempty"`

	FavoritesConfigurations *FavoritesConfigurations `json:"favorites_configurations,omitempty"`

	ContactsConfigurations *ContactsConfigurations `json:"contacts_configurations,omitempty"`

	DownloadsConfigurations *DownloadsConfigurations `json:"downloads_configurations,omitempty"`

	LinksConfigurations *LinksConfigurations `json:"links_configurations,omitempty"`

	SearchesConfigurations *SearchesConfigurations `json:"searches_configurations,omitempty"`

	SavedGamesConfigurations *SavedGamesConfigurations `json:"saved_games_configurations,omitempty"`
}

func (o FolderRedirectionV2Options) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FolderRedirectionV2Options struct{}"
	}

	return strings.Join([]string{"FolderRedirectionV2Options", string(data)}, " ")
}
