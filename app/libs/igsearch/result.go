package igsearch

// IGSearchResult ...
type IGSearchResult struct {
	HasMode  bool              `json:"has_more,emitempty"`
	Hashtags []IGSearchHashtag `json:"hashtags,emitempty"`
	Places   []IGSearchPlace   `json:"places,emitempty"`
	Status   string            `json:"status,emitempty"`
	Users    []IGSearchUser    `json:"users,emitempty"`
}

// IGSearchHashtag ...
type IGSearchHashtag struct {
	Position int64                 `json:"position,emitempty"`
	Hashtag  IGSearchHashtagDetail `json:"hashtag,emitempty"`
}

// IGSearchHashtagDetail ...
type IGSearchHashtagDetail struct {
	ID         int64  `json:"id,emitempty"`
	MediaCount int64  `json:"media_count,emitempty"`
	Name       string `json:"name,emitempty"`

	/*
	  "id": 17841562981104915,
	  "media_count": 340,
	  "name": "jmworks"
	*/
}

// IGSearchPlace ...
type IGSearchPlace struct {
	Position int64               `json:"position,emitempty"`
	Place    IGSearchPlaceDetail `json:"place,emitempty"`
}

// IGSearchPlaceDetail ...
type IGSearchPlaceDetail struct {
	Location IGSearchPlaceLocation `json:"location,emitempty"`
	//"media_bundles": [],
	Slug     string `json:"slug,emitempty"`
	Subtitle string `json:"subtitle,emitempty"`
	Title    string `json:"title,emitempty"`

	/*
	  "media_bundles": [],
	  "slug": "johnson-machine-works",
	  "subtitle": "318 N 11th St, Chariton, Iowa",
	  "title": "Johnson Machine Works"
	*/
}

// IGSearchPlaceLocation ...
type IGSearchPlaceLocation struct {
	Address        string `json:"address,emitempty"`
	City           string `json:"city,emitempty"`
	ExternalSource string `json:"external_source,emitempty"`
	//FacebookPlacesID string  `json:"facebook_places_id,emitempty"`
	FacebookPlacesID int64   `json:"facebook_places_id,emitempty"` // it must be int64, it's not string.
	Lat              float64 `json:"lat,emitempty"`
	Lng              float64 `json:"lng,emitempty"`
	Name             string  `json:"name,emitempty"`
	Pk               string  `json:"pk,emitempty"`

	/*
	   "address": "318 N 11th St",
	   "city": "Chariton, Iowa",
	   "external_source": "facebook_places",
	   "facebook_places_id": 1425020417814317,
	   "lat": 41.01748,
	   "lng": -93.30976,
	   "name": "Johnson Machine Works",
	   "pk": "1029303448"
	*/
}

// IGSearchUser ...
type IGSearchUser struct {
	Position int64              `json:"position,emitempty"`
	User     IGSearchUserDetail `json:"user,emitempty"`
}

// IGSearchUserDetail ...
type IGSearchUserDetail struct {
	Byline        string `json:"by_line,emitempty"`
	FollowerCount int64  `json:"follower_count,emitempty"`
	Following     bool   `json:"following,emitempty"`

	FullName                   string `json:"full_name,emitempty"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture,emitempty"`
	IsPrivate                  bool   `json:"is_private,emitempty"`
	IsVerified                 bool   `json:"is_verified,emitempty"`

	MutualFollowersCount float64 `json:"mutual_followers_count,emitempty"`
	OutgoingRequest      bool    `json:"outgoing_request,emitempty"`
	Pk                   string  `json:"pk,emitempty"`
	ProfilePicID         string  `json:"profile_pic_id,emitempty"`
	ProfilePicURL        string  `json:"profile_pic_url,emitempty"`
	UnseenCount          int64   `json:"unseen_count,emitempty"`
	Username             string  `json:"username,emitempty"`

	/*
	  "byline": "18.5k followers",
	  "follower_count": 18519,
	  "following": false,
	  "full_name": "Joe Mio / ミオジョウ",
	  "has_anonymous_profile_picture": false,
	  "is_private": false,
	  "is_verified": false,
	  "mutual_followers_count": 0.0,
	  "outgoing_request": false,
	  "pk": "118181",
	  "profile_pic_id": "1293589924259512202_118181",
	  "profile_pic_url": "https://scontent-nrt1-1.cdninstagram.com/t51.2885-19/s150x150/13652205_1168813256493819_1560883597_a.jpg",
	  "unseen_count": 0,
	  "username": "jmworks"
	*/
}
