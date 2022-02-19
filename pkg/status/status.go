package status

import "fmt"

func Banner() {
	// http://patorjk.com/software/taag/#p=display&f=Small%20Slant&t=Wiremock
	var banner = `
  _      ___                        __  
 | | /| / (_)______ __ _  ___  ____/ /__
 | |/ |/ / / __/ -_)  ' \/ _ \/ __/  '_/
 |__/|__/_/_/  \__/_/_/_/\___/\__/_/\_\
`
	fmt.Printf("%s", banner)
}

func Pattern() string {
	pattern := `
Wiremock require pattern: 
project
└── mock
   ├── login
   │   └── route.yml
   └── user
       ├── response
       │   └── user.json
       └── route.yml

Please back to root project.
`
	return pattern
}

func Started(port string) {
	started := fmt.Sprintf(`
-> wiremock server started on %s
`, port)
	fmt.Println(started)
}
