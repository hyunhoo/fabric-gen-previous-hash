Generate PreviousHash in Hyperledger Fabric
======================

# 1. Explanation
Block header contain three fields. Which is Number, DataHash and PreviousHash.

According to [**readthedocs**](https://hyperledger-fabric.readthedocs.io), PreviousHash is The hash from the previous block header.

Put previous block's Number, DataHash and PreviousHash then you can compare output with current block's PreviousHash 

# 2. How to use
<pre><code>$ go build -o generator
$ ./generator {Number} {DataHash} {PreviousHash}
</code></pre>
