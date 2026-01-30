package discord

import (
    "ProtoDiscord/commands"
    "ProtoDiscord/core"
    "ProtoDiscord/messages"
    "log/slog"

    "github.com/bwmarrin/discordgo"
)

func onReady(_ *discordgo.Session, r *discordgo.Ready) {
    core.System.Logger().Info("Discord connection established")
    core.System.Logger().Info("Logged in", slog.String("as", r.User.String()))
    core.Spawn(&ApplicationCommands{}, core.ApplicationCommandsPID.Id)
    core.System.Root.Send(core.ApplicationCommandsPID, &messages.Register{})
}

func onCommand(_ *discordgo.Session, i *discordgo.InteractionCreate) {
    if i.Type != discordgo.InteractionApplicationCommand {
        return
    }

    data := i.ApplicationCommandData()
    core.System.Logger().Info("Received command", slog.String("command", data.Name))

    for cmd := range commands.Map {
        if data.Name == commands.Map[cmd].Name {
            actorName := "command-" + data.Name
            pid := core.SpawnCommand(cmd, actorName)
            core.System.Root.Send(pid, &messages.ExecuteCommand{Interaction: i.Interaction, CommandData: &data})
            return
        }
    }
}
