# chess

Console-based, two-player chess in Golang.

## UI

```
$ ./chess

   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |
 - +-----+-----+-----+-----+-----+-----+-----+-----+
 0 | ● R | ● K | ● B | ● Q | ● G | ● B | ● K | ● R |
 1 | ● P | ● P | ● P | ● P | ● P | ● P | ● P | ● P |
 2 |     |     |     |     |     |     |     |     |
 3 |     |     |     |     |     |     |     |     |
 4 |     |     |     |     |     |     |     |     |
 5 |     |     |     |     |     |     |     |     |
 6 | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P |
 7 | ○ R | ○ K | ○ B | ○ Q | ○ G | ○ B | ○ K | ○ R |
 - +-----+-----+-----+-----+-----+-----+-----+-----+

WHITE plays. Enter next ○ move: e7 e5

   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |
 - +-----+-----+-----+-----+-----+-----+-----+-----+
 0 | ● R | ● K | ● B | ● Q | ● G | ● B | ● K | ● R |
 1 | ● P | ● P | ● P | ● P | ● P | ● P | ● P | ● P |
 2 |     |     |     |     |     |     |     |     |
 3 |     |     |     |     |     |     |     |     |
 4 |     |     |     |     | ○ P |     |     |     |
 5 |     |     |     |     |     |     |     |     |
 6 | ○ P | ○ P | ○ P | ○ P |     | ○ P | ○ P | ○ P |
 7 | ○ R | ○ K | ○ B | ○ Q | ○ G | ○ B | ○ K | ○ R |
 - +-----+-----+-----+-----+-----+-----+-----+-----+

BLACK plays. Enter next ● move:
```

## Run

```
$ go run .
```

## Test

```
$ go test
```

## Implementation

* [x] Pieces movement
* [x] Capturing
* [X] Check
* [X] Checkmate
* [X] Resignation
* [ ] Promotion
* [ ] Castling
* [ ] Stalemate
* [ ] En passant


## License

[MIT](LICENSE)
