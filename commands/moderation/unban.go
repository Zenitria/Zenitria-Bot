package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleUnban(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	id := data.Options[0].StringValue()

	s.GuildBanDelete(i.GuildID, id)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🚷・%s has been unbanned", id),
		Description: "🛡️・**Moderator**: " + i.Member.User.Mention(),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/256nKc4aH94AAAAd/pls-unban-me-unban-me.gif",
		},
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	}

	s.InteractionRespond(i.Interaction, response)
}
