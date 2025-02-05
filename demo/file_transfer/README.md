# File transfer demo

This demo transfers a file between two SCION ASes and shows how throughput can be improved by load
balancing the traffic among two non-overlapping paths.

To run the demo:

1. [set up the development environment](https://docs.scion.org/en/latest/build/setup.html)
1. `bazel test --test_output=streamed --cache_test_results=no //demo/file_transfer:file_transfer`

The topology consists of two ASes connected by two links, each thottled to 16 mbps of throughput:

```text
+---------------+  #1               #1  +---------------+
|               | ------throttling----- |               |
| AS ff00:0:110 |                       | AS ff00:0:111 |
|               | ------throttling----- |               |
+---------------+  #2               #2  +---------------+
```

The file is transferred from AS ff00:0:111 to AS ff00:0:110.

The SCION-IP gateway on the sending side is configured to use either one or two paths,
respectively.

The off-the-shelf file transfer application ([bbcp](https://github.com/eeertekin/bbcp))
opens several TCP connections to transfer the data. Those connections are then spread
by the gateway among the two paths.

