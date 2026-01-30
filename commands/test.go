package commands

import (
    "ProtoDiscord/core"
    "ProtoDiscord/messages"
    "log/slog"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/bwmarrin/discordgo"
)

type Test struct{}

func (e Test) Receive(ctx actor.Context) {
    switch msg := ctx.Message().(type) {
    case *actor.Started:
        ctx.Logger().Info("Started, actor started", slog.String("actor", ctx.Self().Id))

    case *actor.Stopped:
        ctx.Logger().Info("Stopped, actor has shut down", slog.String("actor", ctx.Self().Id))

    case *messages.ExecuteCommand:
        ctx.Send(core.DiscordPID, &messages.Respond{Interaction: msg.Interaction, Response: &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "Receiving you loud and clear " + msg.Interaction.Member.DisplayName(),
                Flags:   discordgo.MessageFlagsEphemeral,
            },
        }})
        ctx.Stop(ctx.Self())
    }
}
