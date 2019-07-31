# Ping

Ping is a network utility allowing to test host reachability through IP.

## Installation

To install ping, you must first:

```bash
$ git clone git@github.com:Geoffrey42/ping.git
$ cd ping/
```

Then build the binary using:

```bash
$ make
```

## Usage

Use it like the already installed ping. Only the following options are supported:

```
usage: ping [-hv] host
```

With host being only IPv4. FQDN is supported concerning the returning paquet but 
there is no DNS resolution.

## Contributing
Pull requests are welcome.
For more details, please refers to our [contributing file](.github/CONTRIBUTING/contributing.md).

## License

[MIT](https://choosealicense.com/licenses/mit/)
