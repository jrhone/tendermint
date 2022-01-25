# Assumptions

All aliens move before we check for fights, on each iteration.
- There is no fighting immediately after initialization.
- Aliens on a city get destroyed along with it, even if not involved in a fight.

Map does not need to be fully complete.
- Cities can be declared as links only, without their own line.
- Links not specified are not dynamically created. One way links are possible.

# Run

go mod download
go run cmd/app/main.go <num_aliens>

# World Map

Located in `configs/world_map.txt`

# Tests

go test ./...
