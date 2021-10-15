# Ambilight

for now this only works on linux, but feel free to [create a PR](https://github.com/Martijn-Faber/ambilight/pulls) to support your platform

[![Demo video](https://img.youtube.com/vi/ftrAhJZCl8g/maxresdefault.jpg)](https://www.youtube.com/watch?v=ftrAhJZCl8g)

## Partslist

- Esp
- [Compatible led strip](#compatible-led-strips)

## installation

1. install WLED on the Esp. You can find great tutorials online or look at their [wiki](https://github.com/Aircoookie/WLED/wiki/)

### Binaries

If you want to install a binary, please take a look at the [releases page](https://github.com/Martijn-Faber/ambilight/releases)

### From source

To install it from source you'll first need to [install Go](https://golang.org/doc/install) and clone the [repo](https://github.com/Martijn-Faber/ambilight.git)

```bash
git clone https://github.com/Martijn-Faber/ambilight.git
```

after that you'll need to rename the `config.env.example` to `config.env` and set all the variables. the next step is to install all the dependencies by running `go get` in your terminal. To build the binary you can run `go build` and to run it you can do `./ambilight` in your terminal.

## Contributing

feel free to contribute to this project
check [CONTRIBUTING](./CONTRIBUTING.md) before you commiting

## Compatible LED Strips

| Type                   | Voltage | Comments                      |
| ---------------------- | ------- | ----------------------------- |
| WS2812B                | 5v      |
| WS2813                 | 5v      |
| SK6812                 | 5v      | RGBW                          |
| APA102                 | 5v      | C/D                           |
| WS2801                 | 5v      | C/D                           |
| LPD8806                | 5v      | C/D                           |
| TM1814                 | 12v     | RGBW                          |
| WS2811                 | 12v     | 3-LED segments                |
| WS2815                 | 12v     |
| GS8208                 | 12v     |
| Analog/non-addressable | any     | Requires additional circuitry |

## Author

[@Martijn-Faber](https://github.com/Martijn-Faber)

## Credits

[@afaber999](https://github.com/afaber999)

## License

this project is licensed under the [MIT license](./LICENSE)
