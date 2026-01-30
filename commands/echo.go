package commands

import (
    "ProtoDiscord/core"
    "ProtoDiscord/messages"
    "log/slog"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/bwmarrin/discordgo"
)

type Echo struct{}

func (e Echo) Receive(ctx actor.Context) {
    switch msg := ctx.Message().(type) {
    case *actor.Started:
        ctx.Logger().Info("Started, actor started", slog.String("actor", ctx.Self().Id))

    case *actor.Stopped:
        ctx.Logger().Info("Stopped, actor has shut down", slog.String("actor", ctx.Self().Id))

    case *messages.ExecuteCommand:
        ctx.Send(core.DiscordPID, &messages.Respond{Interaction: msg.Interaction, Response: &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "Oh my, it is so echoy in here! " + msg.CommandData.GetOption("message").StringValue(),
                Flags:   discordgo.MessageFlagsEphemeral,
            },
        }})
        ctx.Stop(ctx.Self())
    }
}
