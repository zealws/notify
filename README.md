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

Start the server with `notify-serve`. You can then 

## Notifications over SSH

Add this to your `~/.ssh/config`:

    Host myhost
      RemoteForward 127.0.0.1:42434 127.0.0.1:42434

You can then create notifications remotely using `notify`.
