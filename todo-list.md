#TODO

* Ways to expose `gopark` and `goready` method without exposing underneath complexities:
    * For some custom use cases like multi-producer single-consumer, traditional channel based approach has some 
    performance limitations due to underneath mutex. Lock-free data-structure have been leveraged for better performance
    in such cases. In golang, we can not directly park/awake a go-routine. Channels hide that complexity.
    **Along with Lock-free DS, custom go-routine's async behaviour is required**. 
    `Polling` will not improve or even perform bad compared to channels.
    E.g. Consider the batching case, 10000 producer go-routines and one consumer-goroutine. In case of buffer/queue is full.
    Consumer go-routine should be given priority. Instead of producer-routines polling for empty bucket, 
    they should be `park`ed. Golang scheduler can schedule consumer routine quickly compared to polling option. \
    Lock-free along with `gopark` and `goready` options, above scenario can be achieved.
    
    
* Effective ways to perform Goroutine Local pattern:
    * Pattern: 
        * Upsert on new goroutine creation --> Needs synchronization among Ps
        * Read/Update within go-routine --> Is atomic operation more than sufficient as no other go-routine will be accessing?
        Does schedule force Processor's cache to sync with MainMemory during go-routine switching?\
        If yes --> even atomic operation is not required \
        If no --> atomic operation is required.
    * Approaches:
        * Use of sharding to reduce contention
        * Using sync.Map(optimized for many reads few write pattern; handles auto-grow, COW), to locate memory location for current goroutine
        * Use of `LinkedList` where value can be `atomic.Value` and map[go-routine-pointer]*Element.

