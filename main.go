import (
    "github.com/bwmarrin/discordgo"
    "github.com/gorilla/feeds"
)

func main() {
    // Créez un nouveau bot Discord en utilisant votre token d'authentification
    bot, err := discordgo.New("TOKEN")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Récupérez le canal Discord où les messages seront publiés
    channel, err := bot.Channel("CHANNEL_ID")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Créez un tableau contenant les URLs des différents flux RSS à suivre
    rssFeeds := []string{
        "http://exemple.com/rss1",
        "http://exemple.com/rss2",
        "http://exemple.com/rss3",
    }

    // Parcourez les URLs des différents flux RSS
    for _, rssFeed := range rssFeeds {
        // Créez un nouvel objet feed pour chaque URL
        feed, err := feeds.Fetch(rssFeed)
        if err != nil {
            fmt.Println(err)
            return
        }

        // Parcourez les articles du flux et publiez leur titre sur Discord
        for _, item := range feed.Items {
            bot.ChannelMessageSend(channel.ID, item.Title)
        }
    }
}
