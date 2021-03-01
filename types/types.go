package types

type TikTokResponse struct {
	StatusCode int `json:"statusCode"`
	ItemList   []struct {
		ID         string `json:"id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"createTime"`
		Video      struct {
			ID           string   `json:"id"`
			Height       int      `json:"height"`
			Width        int      `json:"width"`
			Duration     int      `json:"duration"`
			Ratio        string   `json:"ratio"`
			Cover        string   `json:"cover"`
			OriginCover  string   `json:"originCover"`
			DynamicCover string   `json:"dynamicCover"`
			PlayAddr     string   `json:"playAddr"`
			DownloadAddr string   `json:"downloadAddr"`
			ShareCover   []string `json:"shareCover"`
			ReflowCover  string   `json:"reflowCover"`
		} `json:"video"`
		Author struct {
			ID             string `json:"id"`
			UniqueID       string `json:"uniqueId"`
			Nickname       string `json:"nickname"`
			AvatarThumb    string `json:"avatarThumb"`
			AvatarMedium   string `json:"avatarMedium"`
			AvatarLarger   string `json:"avatarLarger"`
			Signature      string `json:"signature"`
			Verified       bool   `json:"verified"`
			SecUID         string `json:"secUid"`
			Secret         bool   `json:"secret"`
			Ftc            bool   `json:"ftc"`
			Relation       int    `json:"relation"`
			OpenFavorite   bool   `json:"openFavorite"`
			CommentSetting int    `json:"commentSetting"`
			DuetSetting    int    `json:"duetSetting"`
			StitchSetting  int    `json:"stitchSetting"`
			PrivateAccount bool   `json:"privateAccount"`
		} `json:"author"`
		Music struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			PlayURL     string `json:"playUrl"`
			CoverThumb  string `json:"coverThumb"`
			CoverMedium string `json:"coverMedium"`
			CoverLarge  string `json:"coverLarge"`
			AuthorName  string `json:"authorName"`
			Original    bool   `json:"original"`
			Duration    int    `json:"duration"`
			Album       string `json:"album"`
		} `json:"music"`
		Challenges []struct {
			ID            string `json:"id"`
			Title         string `json:"title"`
			Desc          string `json:"desc"`
			ProfileThumb  string `json:"profileThumb"`
			ProfileMedium string `json:"profileMedium"`
			ProfileLarger string `json:"profileLarger"`
			CoverThumb    string `json:"coverThumb"`
			CoverMedium   string `json:"coverMedium"`
			CoverLarger   string `json:"coverLarger"`
			IsCommerce    bool   `json:"isCommerce"`
		} `json:"challenges,omitempty"`
		Stats struct {
			DiggCount    int `json:"diggCount"`
			ShareCount   int `json:"shareCount"`
			CommentCount int `json:"commentCount"`
			PlayCount    int `json:"playCount"`
		} `json:"stats"`
		DuetInfo struct {
			DuetFromID string `json:"duetFromId"`
		} `json:"duetInfo"`
		OriginalItem bool `json:"originalItem"`
		OfficalItem  bool `json:"officalItem"`
		TextExtra    []struct {
			AwemeID      string `json:"awemeId"`
			Start        int    `json:"start"`
			End          int    `json:"end"`
			HashtagName  string `json:"hashtagName"`
			HashtagID    string `json:"hashtagId"`
			Type         int    `json:"type"`
			UserID       string `json:"userId"`
			IsCommerce   bool   `json:"isCommerce"`
			UserUniqueID string `json:"userUniqueId"`
			SecUID       string `json:"secUid"`
		} `json:"textExtra,omitempty"`
		Secret            bool `json:"secret"`
		ForFriend         bool `json:"forFriend"`
		Digged            bool `json:"digged"`
		ItemCommentStatus int  `json:"itemCommentStatus"`
		ShowNotPass       bool `json:"showNotPass"`
		Vl1               bool `json:"vl1"`
		ItemMute          bool `json:"itemMute"`
		AuthorStats       struct {
			FollowingCount int `json:"followingCount"`
			FollowerCount  int `json:"followerCount"`
			HeartCount     int `json:"heartCount"`
			VideoCount     int `json:"videoCount"`
			DiggCount      int `json:"diggCount"`
			Heart          int `json:"heart"`
		} `json:"authorStats"`
		PrivateItem    bool `json:"privateItem"`
		DuetEnabled    bool `json:"duetEnabled"`
		StitchEnabled  bool `json:"stitchEnabled"`
		ShareEnabled   bool `json:"shareEnabled"`
		StickersOnItem []struct {
			StickerType int      `json:"stickerType"`
			StickerText []string `json:"stickerText"`
		} `json:"stickersOnItem,omitempty"`
		IsAd bool `json:"isAd"`
	} `json:"itemList"`
	HasMore bool `json:"hasMore"`
}
