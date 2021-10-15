# Ambilight

for now this only works on linux, but feel free to [create a PR](https://github.com/Martijn-Faber/ambilight/pulls) to support your platform

[![Demo video](https://img.youtube.com/vi/ftrAhJZCl8g/hqdefault.jpg)](https://www.youtube.com/watch?v=ftrAhJZCl8g)

## Partslist

- Esp
- [Compatible led strip](#compatible-led-strips)

## installation

1. install WLED on the Esp. You can find great tutorials online or look at their [wiki](https://github.com/Aircoookie/WLED/wiki/)
2. [install Go](https://golang.org/doc/install)
3. rename the `config.env.example` to `config.env`
4. set all the variables
5. install all the dependencies
   ```bash
   go get
   ```
6. build it
   ```bash
   go build
   ```
7. run it
   ```bash
   ./ambilight
   ```

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
