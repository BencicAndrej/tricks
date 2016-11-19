# Clock

Clock is a simple package that wraps the `time.Now()` method from the standard library into an interface.

The reason for this is to explicitly define time as a dependency when used and to enable providing fake 
implementations of time for testing purposes.