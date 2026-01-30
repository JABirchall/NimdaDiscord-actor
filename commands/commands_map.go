package commands

import (
    "github.com/asynkron/protoactor-go/actor"
    "github.com/bwmarrin/discordgo"
)

var (
    Map = map[actor.Actor]*discordgo.ApplicationCommand{
        Echo{}: {
            Name:        "echo",
            Description: "Say something through a bot",
            Options: []*discordgo.ApplicationCommandOption{
                {
                    Name:        "message",
                    Description: "Contents of the message",
                    Type:        discordgo.ApplicationCommandOptionString,
                    Required:    true,
                },
                {
                    Name:        "author",
                    Description: "Whether to prepend message's author",
                    Type:        discordgo.ApplicationCommandOptionBoolean,
                },
            },
        },
        Test{}: {
            Name:        "test",
            Description: "Test command to check bot responsiveness",
        },
    }
)
