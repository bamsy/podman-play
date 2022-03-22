# podman-play

A cli tool to give the podman play kube command docker-compose commands while still using the kubernetes yaml files for pod setup.

## Purpose

I was using podman play kube quite a bit and was getting tired of the verbose starting, stopping of pods. I chose similar commands to docker-compose because that is what I used in the past.

## Install

NOTE: podman 3 is required to be installed on the system for this to work.

```
git clone github.com/bamsy/podman-play
cd podman-play
go build .
```

## Example

NOTE: All commands currently use kube.yaml as a default yaml file for podman play kube. An optional parameter can be passed in with a path to the yaml file.

### Starting a pod

NOTE: Currently the log driver is defaulted to journald.

```
podman-play up
```

### Stopping a pod

```
podman-play down
```

### Pulling images in a pod yaml file
```
podman-play pull
```

## What is missing currently?

- Need to allow configuration of what log driver to pass in. It is currently defaulted to journald.
- Need to setup a ConfigMap default and look for it on the up command. This will hopefully work similar to using a .env with docker-compose


