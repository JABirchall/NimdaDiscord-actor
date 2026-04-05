package core

import (
    "github.com/asynkron/protoactor-go/actor"
)

var (
    System                 *actor.ActorSystem
    DiscordPID             = &actor.PID{Address: "local", Id: "Discord"}
    ApplicationCommandsPID = &actor.PID{Address: "local", Id: "ApplicationCommands"}
)

func Spawn(ref actor.Actor, name string) {
    _, err := System.Root.SpawnNamed(actor.PropsFromProducer(func() actor.Actor { return ref }), name)
    if err != nil {
        panic(err)
    }
}

// ActorFactory can be implemented by stateful command actors to produce a fresh
// instance for each spawn instead of sharing the map-key instance.
type ActorFactory interface {
    NewActor() actor.Actor
}

func SpawnCommand(ref actor.Actor, name string) *actor.PID {
    var producer func() actor.Actor
    if factory, ok := ref.(ActorFactory); ok {
        producer = factory.NewActor
    } else {
        producer = func() actor.Actor { return ref }
    }
    return System.Root.SpawnPrefix(actor.PropsFromProducer(producer), name)
}
