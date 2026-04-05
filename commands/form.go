package commands

import (
    "ProtoDiscord/core"
    "ProtoDiscord/messages"
    "log/slog"
    "time"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/bwmarrin/discordgo"
)

const formModalPrefix = "form-modal:"

type Form struct {
    interaction *discordgo.Interaction
}

func (f *Form) NewActor() actor.Actor { return &Form{} }

func (f *Form) Receive(ctx actor.Context) {
    switch msg := ctx.Message().(type) {
    case *actor.Started:
        ctx.Logger().Info("Started, actor started", slog.String("actor", ctx.Self().Id))

    case *actor.Stopped:
        ctx.Logger().Info("Stopped, actor has shut down", slog.String("actor", ctx.Self().Id))

    case *messages.ExecuteCommand:
        f.interaction = msg.Interaction
        modalID := formModalPrefix + msg.Interaction.ID
        ctx.Send(core.DiscordPID, &messages.Respond{
            Interaction: msg.Interaction,
            Response: &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseModal,
                Data: &discordgo.InteractionResponseData{
                    CustomID: modalID,
                    Title:    "Tell us about yourself",
                    Components: []discordgo.MessageComponent{
                        discordgo.ActionsRow{
                            Components: []discordgo.MessageComponent{
                                discordgo.TextInput{
                                    CustomID:    "name",
                                    Label:       "Name",
                                    Style:       discordgo.TextInputShort,
                                    Placeholder: "Your name",
                                    Required:    true,
                                },
                            },
                        },
                        discordgo.ActionsRow{
                            Components: []discordgo.MessageComponent{
                                discordgo.TextInput{
                                    CustomID:    "bio",
                                    Label:       "Bio",
                                    Style:       discordgo.TextInputParagraph,
                                    Placeholder: "A short bio",
                                    Required:    false,
                                },
                            },
                        },
                    },
                },
            },
        })

        ctx.Send(core.DiscordPID, &messages.RegisterInteractionHandler{
            CustomID: modalID,
            Handler:  ctx.Self(),
        })

        ctx.SetReceiveTimeout(1 * time.Minute)

    case *actor.ReceiveTimeout:
        ctx.Logger().Info("Form timed out, stopping actor", slog.String("actor", ctx.Self().Id))
        ctx.Send(core.DiscordPID, &messages.FollowUp{
            Interaction: f.interaction,
            Content:     "The form has expired. Please run `/form` again if you'd like to submit.",
            Flags:       discordgo.MessageFlagsEphemeral,
        })
        core.DeleteInteraction(formModalPrefix + f.interaction.ID)
        ctx.Stop(ctx.Self())

    case *messages.InteractionEvent:
        data := msg.Interaction.ModalSubmitData()

        name := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
        bio := data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

        content := "Thanks, **" + name + "**!"
        if bio != "" {
            content += "\n> " + bio
        }

        ctx.Send(core.DiscordPID, &messages.Respond{
            Interaction: msg.Interaction,
            Response: &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseChannelMessageWithSource,
                Data: &discordgo.InteractionResponseData{
                    Content: content,
                    Flags:   discordgo.MessageFlagsEphemeral,
                },
            },
        })

        ctx.Stop(ctx.Self())
    }
}
