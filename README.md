# go-tsort

UNIX tsort utility written in go

## Dependencies

* Go >= 1.16 (build time)

## Note

Output of `tsort` may be different between various implementations. It's
because POSIX unclear about which algorithm shall be used to process
dependency graph. Conforming applications cannot rely on this behavior.
