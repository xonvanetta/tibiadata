package v2

import (
	"context"
	"fmt"
)

type NewsResponse struct {
	News        News         `json:"news"`
	Information *Information `json:"information"`
}

type News struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    *Timezone `json:"date"`
}

func (c client) News(context context.Context, newsId int) (*NewsResponse, error) {
	newsResponse := &NewsResponse{}
	url := tibiaDataURL(fmt.Sprintf("news/%d.json", newsId))
	err := c.client.Get(context, url, newsResponse)
	return newsResponse, err
}
