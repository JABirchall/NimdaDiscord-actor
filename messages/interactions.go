package messages

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/bwmarrin/discordgo"
)

// RegisterInteractionHandler registers a temporary actor to handle a follow-up
// interaction with the given custom ID. Send this to DiscordPID from a command actor.
type RegisterInteractionHandler struct {
	CustomID string
	Handler  *actor.PID
}

// InteractionEvent is dispatched to a registered handler when a matching
// component or modal interaction is received from Discord.
type InteractionEvent struct {
	Interaction *discordgo.Interaction
}
