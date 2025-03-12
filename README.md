**Build the binary**:
    
    `cd client go build -o eth-client`
    
- **Run with default values**:
    
    
    `./eth-client`
    
- **Run with a specific input file**:
    
    
    `./eth-client -file="../input.json"`
    
- **Pipe JSON input directly**:
    
    
    `cat ../input.json | ./eth-client`
    
    or
    
    `echo '{"jsonrpc":"2.0","method":"eth_blockNumber","id":2}' | ./eth-client`


## How To run Tests
1. To run the tests, navigate to the client directory and run:
    `go test -v`




## What could be added to make it better
1. Auth token header can be added so that, we can call other RPC endpoints which require a Bearer Token to Authentication
2. Method based cases can be added, like for each method namely `eth_blockNumber` and `eth_getBlockByNumber` we can give method based input instead of providing whole json as input.
3. Chaining of the result for the two method calls to Polygon RPC i.e calling `eth_blockNumber` should automatically give a (Y/n) prompt to call `eth_getBlockByNumber` for the blockNumber received in response from `eth_blockNumber` 