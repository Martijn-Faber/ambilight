# Ambilight

for now this only works on linux, but feel free to [create a PR](/pulls) to support your platform

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
6. run it
   ```bash
   go run main.go
   ```

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
