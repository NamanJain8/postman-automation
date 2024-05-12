# postman-automation
Demo code for postman automation using local server

## How to Run

1. Start token server by executing `go run .` from `./token-server`
2. Start catalogue server by executing `go run .` from `./catalogue-service`
3. Import postman collection and play around by changing USERNAME collection variable.

## Set up daemon

1. Move the plist file to `~/Library/LanchAgents/`
2. Load the daemon service `launchctl load ~/Library/LaunchAgents/com.local.token.plist`
3. Validate using: `curl "http://localhost:4040/token" --data-raw '{"username":"john"}'`
