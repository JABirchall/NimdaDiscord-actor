package discord

import (
    "ProtoDiscord/messages"
    "log/slog"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/bwmarrin/discordgo"
)

type Connect struct {
    appid   string
    discord *discordgo.Session
}

func (dc *Connect) Receive(ctx actor.Context) {
    var err error
    switch msg := ctx.Message().(type) {
    case *actor.Started:
        ctx.Logger().Info("Started, actor started", slog.String("actor", ctx.Self().Id))

    case *actor.Stopped:
        ctx.Logger().Info("Stopped, actor has shut down", slog.String("actor", ctx.Self().Id))

    case *messages.Connect:
        dc.appid = msg.App

        if dc.discord, err = discordgo.New("Bot " + msg.Token); err != nil {
            ctx.Logger().Error("Discord connection failed", slog.String("error", err.Error()))
            return
        }

        dc.discord.AddHandler(onReady)
        dc.discord.AddHandler(onCommand)

        if err = dc.discord.Open(); err != nil {
            ctx.Logger().Error("Discord connection failed", slog.String("error", err.Error()))
            return
        }

    case *messages.Respond:
        if err = dc.discord.InteractionRespond(msg.Interaction, msg.Response); err != nil {
            ctx.Logger().Error("Failed to respond to interaction", slog.String("error", err.Error()))
            return
        }

    case *messages.CommandOverride:
        if _, err = dc.discord.ApplicationCommandBulkOverwrite(dc.appid, "", msg.Commands); err != nil {
            ctx.Logger().Error("Failed to register commands", slog.String("error", err.Error()))
            return
        }

    case *messages.CommandCreate:
        command, err := dc.discord.ApplicationCommandCreate(dc.appid, "", msg.Command)
        if err != nil {
            ctx.Logger().Error("Failed to create command", slog.String("error", err.Error()))
            return
        }

        ctx.Logger().Info("Created command", slog.String("command", command.Name), slog.String("id", command.ID))

        if msg.Delete {
            if err := dc.discord.ApplicationCommandDelete(dc.appid, "", command.ID); err != nil {
                ctx.Logger().Error("Failed to delete command", slog.String("error", err.Error()))
                return
            }
            ctx.Logger().Info("Deleted command", slog.String("command", command.Name), slog.String("id", command.ID))
        }
    }
}
