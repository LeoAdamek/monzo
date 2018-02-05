package monzo


// FeedEntry represents an entry into the user's account feed
type FeedEntry struct {
    AccountID string `json:"account_id"`
    Type string `json:"type"`
    URL string `json:"url"`
    
    Properties struct {
        Title string `json:"title"`
        Body string `json:"body"`
        ImageURL string `json:"image_url"`
        BackgroundColor string `json:"background_color"`
        BodyColor string `json:"body_color"`
        TitleColor string `json:"title"`
    } `json:"params"`
}

// AddFeedItem adds a new item to the user's account feed
func AddFeedItem(item FeedEntry) error {
    

    return nil
}