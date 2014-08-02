package fakku

import (
	"fmt"
)

type ForumCategoriesApiFunction struct{}

func (c ForumCategoriesApiFunction) ConstructApiFunction() string {
	return fmt.Sprintf("%s/forums", ApiHeader)
}

func GetForumCategories() (*ForumCategories, error) {
	var c ForumCategories
	url := ForumCategoriesApiFunction{}
	if err := ApiCall(url, &c); err != nil {
		return nil, err
	} else {
		return &c, nil
	}
}

type ForumCategories struct {
	Categories []*ForumCategory `json:"categories"`
}
type ForumCategory struct {
	Title  string   `json:"category_title"`
	Order  string   `json:"category_order"`
	Forums []*Forum `json:"forums"`
}

type Forum struct {
	Name        string `json:"forum_name"`
	Description string `json:"forum_description"`
	Url         string `json:"forum_url"`
	Posts       uint   `json:"forum_posts"`
	Topics      uint   `json:"forum_topics"`
	Silent      uint   `json:"forum_silent"`
	RecentTopic *Topic `json:"forum_recent_topic"`
	// There are more....
}

type Topic struct {
	Title       string `json:"topic_title"`
	Url         string `json:"topic_url"`
	Time        uint   `json:"topic_time"`
	FirstPostId uint   `json:"topic_first_post_id"`
	LastPostId  uint   `json:"topic_last_post_id"`
	FrontPage   uint   `json:"front_page"`
	Status      uint   `json:"topic_status"`
	Vote        uint   `json:"topic_vote"`
	Type        uint   `json:"topic_type"`
	Poster      string `json:"topic_poster"`
	PosterUrl   string `json:"topic_poster_url"`
}
