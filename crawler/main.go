package main

import (
	"fmt"
	"time"
	"imooc.com/ccmouse/learngo/crawler/engine"
	"imooc.com/ccmouse/learngo/crawler/scheduler"
)

func main() {
		itemChan, err := persist.ItemSaver(
			config.ElasticIndex)
		if err != nil {
			panic(err)
		}
		e := engine.ConcurrentEngine{
			Scheduler:        &scheduler.QueuedScheduler{},
			WorkerCount:      100,
			ItemChan:         itemChan,
			RequestProcessor: engine.Worker,
		}

		e.Run(engine.Request{
			Url: "http://www.zhenai.com/zhenghun",
			Parser: engine.NewFuncParser(
				parser.ParseCityList,
				config.ParseCityList),
		})
	var j chan interface{}
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j <- tick
		fmt.Println(j)
	}
	//launch()

}
