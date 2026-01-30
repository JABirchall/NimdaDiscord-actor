package main

import (
    "ProtoDiscord/core"
    "ProtoDiscord/discord"
    "ProtoDiscord/messages"
    "log"
    "log/slog"
    "os"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/joho/godotenv"
    "github.com/lmittmann/tint"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    core.System = actor.NewActorSystemWithConfig(&actor.Config{LoggerFactory: func(system *actor.ActorSystem) *slog.Logger {
        w := os.Stderr
        // create a new logger
        return slog.New(tint.NewHandler(w, &tint.Options{
            Level: slog.LevelInfo,
        }))
    }})
    core.System.ProcessRegistry.Address = "local"

    core.Spawn(&discord.Connect{}, core.DiscordPID.Id)
    core.System.Root.Send(core.DiscordPID, &messages.Connect{Token: os.Getenv("DISCORD_TOKEN"), App: os.Getenv("DISCORD_APP_ID")})

    event := make(chan os.Signal, 1)
    <-event
    core.System.Root.Stop(core.DiscordPID)
}
