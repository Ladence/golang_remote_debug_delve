# Tutorial: How to debug your Go-service using delve debugger & docker

## Summary

This repo is just a tutorial for my colleagues and students introducing how to debug
go service using delve (for this example i use only docker configuration, but it can be applied on Kubernetes nodes too)

NOTE: go code in main.go is only for lulz, don't do that in your production code :)

## Steps

1. Prepare your service (or you can use my example code in main.go)
2. Prepare Docker multi-staged configuration with 2 steps (watch Dockerfile) :

- Build & install delve debugger. Note: we're building without strip options (-s in ldflags) to keep our binary debug information. -N is used also for turning off optimization. -L is used for showing complete file path in error messages.
- Copy from first stage binary & run delve (which will run our program when any external connection with delve will be established).

Delve is running under specified configuration, we're using flags:

* listen - to specify port under what we want to listen our delve-clients
* headless
* accept-multiclient
* api-version=2
* exec %PATH to a program binary%

Congratulations, now you can run your program in a docker by:

`docker run --security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE`

Note: don't forget about ports mapping. In my case, I need to add `8000:8000 40000:40000` (first port is for service, second one for delve listened port)


## Instruction for GoLand users 

if you like to use IDE GoLand, you should create 2 run-configurations: one for docker and one for remote-debugging (GoLand have a separate one configuration for remote-debug)

Start debug session by starting docker firstly.