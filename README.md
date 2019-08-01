# go_ping

**go_ping** is an implementation of ```ping``` written in Go. go_ping is a network utility allowing to test host reachability through IP.

## Installation

To install go_ping, you must first:

```bash
$ git clone git@github.com:Geoffrey42/go_ping.git
$ cd go_ping/
```

Then build the binary using:

```bash
$ make
```

## Usage

Use it like **ping** although only the following options are supported:

```
usage: go_ping [-hv] host
```

With host being only IPv4. FQDN is supported concerning the returning paquet but 
there is no DNS resolution.

## Contributing
Pull requests are welcome.
For more details, please refers to our [contributing file](.github/CONTRIBUTING/contributing.md).

## Disclaimer
This subject is based on 42 school ft_ping project [here](assets/ft_ping.fr.pdf)
written in C.

## License

[MIT](https://choosealicense.com/licenses/mit/)
