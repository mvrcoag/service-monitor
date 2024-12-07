# Service Monitor

Service Monitor is a simple tool to monitor the status of services on a server. It is designed to be used in a server environment where you have multiple services running and you want to monitor their status.

## Usage

You can execute the program by running the following command:

```bash
go run main.go
```

This will give you a list of available commands:

```bash
Usage:
   [command]

Available Commands:
  add         Add a URL to the system monitor
  completion  Generate the autocompletion script for the specified shell
  drop        Drop a service by index
  help        Help about any command
  list        List all the services
  report      Give a report of the registered URLs

Flags:
  -h, --help   help for this command

Use "[command] --help" for more information about a command.
```

You can add a service to the monitor by running the following command:

```bash
go run main.go add http://example.com
```

This will add the service at the given URL to the monitor. You can then list all the services by running the following command:

```bash
go run main.go list
```

This will give you a list of all the services that are being monitored along with their status.

You can also generate a report of the services by running the following command:

```bash
go run main.go report
```

This will give you a report of all the services that are being monitored along with their status.

You can drop a service from the monitor by running the following command:

```bash
go run main.go drop 1
```

This will drop the service at index 1 from the monitor.

## Binary build

You can also build a binary of the program by running the following command:

```bash
go build -o service-monitor main.go
```

This will create a binary called `service-monitor` which you can run directly:

```bash
./service-monitor
```

## Install binary globally (optional)

You can also install the binary globally by running the following command:

```bash
sudo mv service-monitor /usr/local/bin
```

This will allow you to run the program from anywhere on your system:

```bash
service-monitor
```
