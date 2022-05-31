module Day5/geecache

go 1.16


require consistenthash v1.0.0-1
require lru v1.0.0-1 //indect

replace lru => ../lru/

replace consistenthash => ./consistenthash/
