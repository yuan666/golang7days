module Day4

go 1.17

require geecache v1.0.0-1

require lru v1.0.0-1 // indirect

replace lru => ./lru/

replace geecache => ./geecache/
