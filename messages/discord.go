package messages

import "github.com/bwmarrin/discordgo"

type Connect struct {
    App   string
    Token string
}

type CommandOverride struct {
    Commands []*discordgo.ApplicationCommand
}

type Respond struct {
    Interaction *discordgo.Interaction
    Response    *discordgo.InteractionResponse
}

type CommandCreate struct {
    Delete  bool
    Command *discordgo.ApplicationCommand
}

type Register struct{}
