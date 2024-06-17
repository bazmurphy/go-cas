# CAS in Go

The Compare-and-Swap (CAS) operation is a key concept in distributed systems, particularly in scenarios where multiple clients may try to modify the same data concurrently. It provides a way to atomically update a value only if it hasn't been modified by another client since the last read.

In this project, the server implements the CAS operation in the `Cas` method. It checks if the `expectedValue` matches the current value associated with the key, and if so, updates it with the `newValue`. This ensures that the update is performed only if the value hasn't been modified by another client concurrently.

The client demonstrates the usage of the CAS operation by performing a series of CAS requests, attempting to increment the value associated with a key. Each client instance tries to perform a specified number of CAS operations, simulating concurrent access to the shared data.

1. Protocol Buffer (`proto`) file:
   - Defines the gRPC service `KeyValueStoreService` and the associated request and response messages for the `Get`, `Set`, and `Cas` operations.
   - Serves as the contract between the server and the client.

2. Server implementation:
   - Implements the `KeyValueStoreService` server using gRPC.
   - Maintains an in-memory key-value store using a `map[string]int64`.
   - Provides implementations for the `Get`, `Set`, and `Cas` operations.
   - Uses a read-write mutex (`sync.RWMutex`) to synchronize access to the key-value store.
   - The `Cas` operation checks if the `expectedValue` matches the current value associated with the key, and if so, updates it with the `newValue`. It returns a boolean indicating the success or failure of the operation.

3. Client implementation:
   - Acts as a client to the `KeyValueStoreService` server.
   - Establishes a connection to the server using gRPC.
   - Provides methods to invoke the `Get`, `Set`, and `Cas` operations on the server.
   - The `main` function demonstrates the usage of the client by setting an initial value and then performing a series of CAS operations.