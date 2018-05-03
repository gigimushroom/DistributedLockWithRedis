## Overview
Distributed locks are a very useful primitive in many environments where different processes must operate with shared resources in a mutually exclusive way.

## Safety property: 
Mutual exclusion. At any given moment, only one client can hold a lock.
## Liveness property A: 
Deadlock free. Eventually it is always possible to acquire a lock, even if the client that locked a resource crashes or gets partitioned.
## Liveness property B: 
Fault tolerance. As long as the majority of Redis nodes are up, clients are able to acquire and release locks.

## Implementation
Before trying to overcome the limitation of the single instance setup described above, let’s check how to do it correctly in this simple case, since this is actually a viable solution in applications where a race condition from time to time is acceptable, and because locking into a single instance is the foundation we’ll use for the distributed algorithm described here.

To acquire the lock, the way to go is the following:
```
SET resource_name my_random_value NX PX 30000
```
The command will set the key only if it does not already exist (NX option), with an expire of 30000 milliseconds (PX option). The key is set to a value “myrandomvalue”. This value must be unique across all clients and all lock requests.

Basically the random value is used in order to release the lock in a safe way, with a script that tells Redis: remove the key only if it exists and the value stored at the key is exactly the one I expect to be. This is accomplished by the following Lua script:

```
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
```