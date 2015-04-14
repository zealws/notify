# notify

Notification command-line utility for generating desktop notifications from a command line.

## Requirements

- `libnotify`

## Usage

    $ notify COMMAND

Example:

    $ notify sleep 3

    $ notify make -j 50 # notify after long-running command

## Client/Server

Start the server with `notify-serve`. You can then send notifications with curl:

    curl http://localhost:42434/ -d 'Notification here!'

Or from a remote host via the `notify` utility. Just set the NOTIFY_ADDR to the `host:port` that the server is bound on:

    NOTIFY_ADDR=1.2.3.4:42434 notify sleep 3

## Notifications over SSH

Add this to your `~/.ssh/config`:

    Host myhost
      RemoteForward 127.0.0.1:42434 127.0.0.1:42434

You can then create notifications remotely using `notify`.
