package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Treasure Hunt Scenario

Imagine you're coordinating a group of treasure hunters who are exploring different areas to find a hidden treasure. Each hunter is assigned a specific area, and they will report back if they find the treasure. The twist is that as soon as one hunter finds the treasure, you need to notify all other hunters to stop their search immediately.

Flow Explanation
Treasure Hunters:

Each hunter independently searches their assigned area. The time it takes for them to search is random; some might finish quickly, while others take longer.
After finishing their search, each hunter reports back whether they found the treasure.
If any hunter finds the treasure, this news needs to reach all the other hunters immediately so they can stop searching.
Communication:

The hunters need a way to communicate their success or failure back to you, the coordinator.
You, as the coordinator, must also be able to tell all hunters to stop if the treasure is found.
The coordination is crucial because you want to minimize unnecessary searching once the treasure is located.
Cancellation:

As soon as one hunter reports finding the treasure, all ongoing searches by other hunters should be halted.
The coordination mechanism should ensure that as soon as the treasure is found, all other hunters receive the stop order.
Graceful Completion:

The whole operation should end smoothly. If a hunter finds the treasure, all should wrap up their tasks.
If the treasure is not found by any hunter, all should report back their completion.
Your Task
Coordinate the Search:

Organize the treasure hunt, ensuring that all hunters start searching their assigned areas simultaneously.
Make sure to handle communication between the hunters and the coordinator.
Manage Cancellation:

Design a way for the hunters to know when they should stop searching if another hunter finds the treasure.
Complete the Operation:

Ensure that the search ends smoothly, with either the treasure found or all hunters reporting back after completing their search.
*/

// Types

// represents participants
type TreasureHunter string
type TreasureHunters []TreasureHunter

// Constants
// treasure found channel
var treasureFoundChan = make(chan TreasureHunter)

var target = 7
var defaultWaitTime = 4

var once sync.Once

func TreasureHunt() {
	var wg sync.WaitGroup

	fmt.Println("Treasure Hunt")
	noOfParticipants := 5
	treasureHunters := make([]TreasureHunter, noOfParticipants)

	for i := 0; i < noOfParticipants; i++ {
		treasureHunters[i] = TreasureHunter(fmt.Sprintf("Treasure Hunter %d", i+1))
	}

	fmt.Println("treasureHunters:", treasureHunters)

	// start go routine for each
	wg.Add(noOfParticipants)

	for _, treasureHunter := range treasureHunters {
		go startHunt(treasureHunter, &wg)
	}

	// wait and analyse hunt
	go func() {
		select {
		case winner := <-treasureFoundChan:
			fmt.Printf("Winner was %s\n", winner)

			// close channel notify all other treasure hunters to exit loop
			close(treasureFoundChan)
			return
		}
	}()

	wg.Wait()
}

func startHunt(treasureHunter TreasureHunter, wg *sync.WaitGroup) {
	defer wg.Done()
	// generate random hunter speed
	hunterSpeed := rand.Intn(2) + 1 // 1 or 2
	fmt.Printf("hunter %v with speed %d\n", treasureHunter, hunterSpeed)

	for {

		select {
		// treasure already found, exit their own search
		case <-treasureFoundChan:
			return

		// search for treasure
		default:
			// generate random number
			treasureAcquired := rand.Intn(10)
			fmt.Printf("hunter %v found treasure %d. Target is still: %d.\n", treasureHunter, treasureAcquired, target)

			if treasureAcquired == target {
				// treasure was found, notify channel
				once.Do(func() { treasureFoundChan <- treasureHunter }) // use once.Do to limit only one execution across all goroutines
				return
			}

			// simulate wait time
			time.Sleep(time.Duration(defaultWaitTime/hunterSpeed) * time.Second)

		}
	}

}
