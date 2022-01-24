# Assumptions

All aliens move before we check for fights, on each iteration.
- There is no fighting immediately after initialization.
- Aliens on a city get destroyed along with it, even if not involved in a fight.

Map needs to be fully complete.
- Each city declared on a line
- Links not specified are not dynamically created. One way links are possible.

# Run

go mod download
go run cmd/app/main.go <num_aliens>

# World Map

Located in `configs/world_map.txt`

# Tests

go test ./...
