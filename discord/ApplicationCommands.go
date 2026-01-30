package discord

import (
    "ProtoDiscord/commands"
    "ProtoDiscord/core"
    "ProtoDiscord/messages"
    "log/slog"

    "github.com/asynkron/protoactor-go/actor"
)

type ApplicationCommands struct{}

func (ac ApplicationCommands) Receive(ctx actor.Context) {
    switch ctx.Message().(type) {
    case *actor.Started:
        ctx.Logger().Info("Started, actor started", slog.String("actor", ctx.Self().Id))

    case *actor.Stopped:
        ctx.Logger().Info("Stopped, actor has shut down", slog.String("actor", ctx.Self().Id))

    case *messages.Register:
        ctx.Logger().Info("Registering application commands", slog.Int("count", len(commands.Map)))

        for _, cmd := range commands.Map {
            ctx.Send(core.DiscordPID, &messages.CommandCreate{Command: cmd, Delete: false})
        }

        ctx.Stop(ctx.Self())
    }
}
