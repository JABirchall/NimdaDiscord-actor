package core

import (
    "sync"

    "github.com/asynkron/protoactor-go/actor"
)

var interactionHandlers sync.Map

// RegisterInteraction stores a one-shot handler PID for a given custom ID.
func RegisterInteraction(customID string, pid *actor.PID) {
    interactionHandlers.Store(customID, pid)
}

// ConsumeInteraction retrieves and removes the handler for a given custom ID.
// Returns false if no handler is registered.
func ConsumeInteraction(customID string) (*actor.PID, bool) {
    val, ok := interactionHandlers.LoadAndDelete(customID)
    if !ok {
        return nil, false
    }
    return val.(*actor.PID), true
}

func DeleteInteraction(customID string) {
    interactionHandlers.Delete(customID)
}
