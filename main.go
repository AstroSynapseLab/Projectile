// .projectile
/*

env/
	local/
		air.toml
		Dockerfile
		docker-compose.yaml
config/
	main.go
	main.yaml
	auth.yaml
main.go

*/

package main

import (
	"github.com/AstroSynapseLab/Projectile/cmd"
)

func main() {
	cmd.Execute()
}
