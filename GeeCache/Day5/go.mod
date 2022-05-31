module Day5

go 1.16


require geecache v1.0.0-1

require consistenthash v1.0.0-1

require lru v1.0.0-1 //indect

replace lru => ./lru/

replace consistenthash => ./geecache/consistenthash/

replace geecache => ./geecache/