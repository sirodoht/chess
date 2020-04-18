# chess

Console-based, two-player chess in Golang.

## UI

```
$ go run .
+---+-----+-----+-----+-----+-----+-----+-----+-----+
|   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |
+---+-----+-----+-----+-----+-----+-----+-----+-----+
| 1 | ● R | ● K | ● B | ● Q | ● G | ● B | ● K | ● R |
| 2 | ● P | ● P | ● P | ● P | ● P | ● P | ● P | ● P |
| 3 |     |     |     |     |     |     |     |     |
| 4 |     |     |     |     |     |     |     |     |
| 5 |     |     |     |     |     |     |     |     |
| 6 |     |     |     |     |     |     |     |     |
| 7 | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P | ○ P |
| 8 | ○ R | ○ K | ○ B | ○ Q | ○ G | ○ B | ○ K | ○ R |
+---+-----+-----+-----+-----+-----+-----+-----+-----+
WHITE plays. Enter next move: e7 e5

MOVE: white ○ Pawn moved to e5
+---+-----+-----+-----+-----+-----+-----+-----+-----+
|   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |
+---+-----+-----+-----+-----+-----+-----+-----+-----+
| 1 | ● R | ● K | ● B | ● Q | ● G | ● B | ● K | ● R |
| 2 | ● P | ● P | ● P | ● P | ● P | ● P | ● P | ● P |
| 3 |     |     |     |     |     |     |     |     |
| 4 |     |     |     |     |     |     |     |     |
| 5 |     |     |     |     | ○ P |     |     |     |
| 6 |     |     |     |     |     |     |     |     |
| 7 | ○ P | ○ P | ○ P | ○ P |     | ○ P | ○ P | ○ P |
| 8 | ○ R | ○ K | ○ B | ○ Q | ○ G | ○ B | ○ K | ○ R |
+---+-----+-----+-----+-----+-----+-----+-----+-----+
BLACK plays. Enter next move:
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

* [x] Pieces moves validation
* [x] Capturing validation
* [ ] Check
* [ ] Checkmate
* [ ] Promotion
* [ ] Castling
* [ ] Stalemate
* [ ] En passant


## License

[MIT](LICENSE)
