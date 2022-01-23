package main

import (
	"context"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var valueLength = []int{10, 20, 50, 100, 200, 1000, 5000, 10000}

const COUNT = 200000

func main() {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, COUNT)
	for i := range keys {
		keys[i] = RandStringRunes(15)
	}
	client := redis.NewClient(&redis.Options{Addr: "192.168.1.186:6388"})
	if err := client.FlushAll(context.Background()).Err(); err != nil {
		panic(err)
	}
	log.Println("| size  | before | after | actual | per | about |")
	log.Println("| :---: | :----: | :---: | :----: | :-: | :---: |")
	for _, size := range valueLength {

		before, err := GetMemory(client)
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			if err := client.Set(context.Background(), key, RandStringRunes(size), 0).Err(); err != nil {
				panic(err)
			}
		}
		after, err := GetMemory(client)
		if err != nil {
			panic(err)
		}
		log.Printf("|%v|%v|%v|%v|%v|%v|", size, before, after, after-before, float64(after-before)/float64(COUNT), math.Round(float64(after-before)/float64(COUNT)))
		if err := client.FlushAll(context.Background()).Err(); err != nil {
			panic(err)
		}
	}

}

func GetMemory(client *redis.Client) (int64, error) {
	c, err := client.Info(context.Background(), "memory").Result()
	if err != nil {
		return 0, err
	}
	// memory info
	//	# Memory
	//used_memory:960152
	//used_memory_human:937.65K
	//used_memory_rss:6705152
	//used_memory_rss_human:6.39M
	//used_memory_peak:22081208
	//used_memory_peak_human:21.06M
	//used_memory_peak_perc:4.35%
	//	used_memory_overhead:912416
	//used_memory_startup:809880
	//used_memory_dataset:47736
	//used_memory_dataset_perc:31.77%
	//	allocator_allocated:988656
	//allocator_active:3309568
	//allocator_resident:6496256
	//total_system_memory:33566478336
	//total_system_memory_human:31.26G
	//used_memory_lua:37888
	//used_memory_lua_human:37.00K
	//used_memory_scripts:0
	//used_memory_scripts_human:0B
	//number_of_cached_scripts:0
	//maxmemory:0
	//maxmemory_human:0B
	//maxmemory_policy:noeviction
	//allocator_frag_ratio:3.35
	//allocator_frag_bytes:2320912
	//allocator_rss_ratio:1.96
	//allocator_rss_bytes:3186688
	//rss_overhead_ratio:1.03
	//rss_overhead_bytes:208896
	//mem_fragmentation_ratio:7.30
	//mem_fragmentation_bytes:5786024
	//mem_not_counted_for_evict:0
	//mem_replication_backlog:0
	//mem_clients_slaves:0
	//mem_clients_normal:102536
	//mem_aof_buffer:0
	//mem_allocator:jemalloc-5.1.0
	//active_defrag_running:0
	//lazyfree_pending_objects:0
	//lazyfreed_objects:0

	memoryDataset := strings.Split(c, "\r\n")[10]
	memoryStr := strings.Split(memoryDataset, ":")[1]
	return strconv.ParseInt(memoryStr, 10, 64)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
