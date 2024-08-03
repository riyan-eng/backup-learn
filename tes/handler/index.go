package handler

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/panjf2000/ants/v2"
)

func Index(c *fiber.Ctx) error {
	defer ants.Release()

	// p, err := ants.NewPool(2, ants.WithOptions(ants.Options{}))
	// if err != nil {
	// 	panic(err)
	// }

	// defer p.Release()

	// var nana int

	// for i := 1; i <= 100000; i++ {
	// 	fmt.Println(i)
	// 	demoFunc()
	// }

	// fmt.Println(nana)

	runTimes := 100000

	var wg sync.WaitGroup
	syncCalculateSum := func(i int) {
		demoFunc(i)
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		i := i
		_ = ants.Submit(func() {
			syncCalculateSum(i)
		})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
	return c.SendString("Hello, World!")
}

func demoFunc(i int) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!: ", i)
}
