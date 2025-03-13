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
4. Writing better terraform with modules for Sandbox and Prod environments.


## Docker/Podman Compile
1. If you have podman just replace `docker` with `podman` in build command, it should work without any hiccups
2. Run `podman build -t polygon-client .`
3. `podman images` - select the imageid for polygon-client
4. Run `podman run -it <imageID>` - ex -------> podman run -it ddb3e1fed564
5. Execute one of the following commands and check the result ---> name of binary in Docker build is `polygon-client`
 
 `polygon-client` ----> runs with default hardcoded value of json

    or specify your own json and pass input to the utility

 `echo '{"jsonrpc":"2.0","method":"eth_blockNumber","id":2}' | polygon-client`


## Terraform
1. cd `./terraform`
2. Run `terraform init --upgrade` 
3. Run `terrform plan` to see the resources.