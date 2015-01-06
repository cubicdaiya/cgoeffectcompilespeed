Comparisons with cgo effect to compile speed for Go.

```bash
$ make
cd go && time go build

real    0m0.167s
user    0m0.127s
sys     0m0.035s
cd cgo && time go build

real    0m0.544s
user    0m0.307s
sys     0m0.218s
cd cgofast && time go build

real    0m0.385s
user    0m0.238s
sys     0m0.133s
$
```
