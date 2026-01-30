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

func SpawnCommand(ref actor.Actor, name string) *actor.PID {
    return System.Root.SpawnPrefix(actor.PropsFromProducer(func() actor.Actor { return ref }), name)
}
