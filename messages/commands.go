package messages

import "github.com/bwmarrin/discordgo"

type ExecuteCommand struct {
    Interaction *discordgo.Interaction
    CommandData *discordgo.ApplicationCommandInteractionData
}
